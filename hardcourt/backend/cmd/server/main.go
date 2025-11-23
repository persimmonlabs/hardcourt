package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/handlers"
	"hardcourt/backend/internal/repository"
	"hardcourt/backend/internal/scraper"
	"hardcourt/backend/internal/scrapers"
	"hardcourt/backend/internal/simulator"
	"hardcourt/backend/internal/websocket"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-redis/redis/v8"
)

func main() {
	// 1. Config
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"
	}

	ctx := context.Background()

	// 2. Database Connection
	db, err := database.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 3. Run Migrations
	if err := db.RunMigrations(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// 4. Initialize Repositories
	matchRepo := repository.NewMatchRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	tournamentRepo := repository.NewTournamentRepository(db)

	// 5. Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	redisCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if _, err := rdb.Ping(redisCtx).Result(); err != nil {
		log.Printf("Warning: Could not connect to Redis at %s: %v. Simulator will run locally only.", redisAddr, err)
	} else {
		log.Println("Connected to Redis")
	}

	// 6. Core Components
	matchUpdateChan := make(chan *domain.Match, 100)
	hub := websocket.NewHub()

	// Initialize scraper aggregator for real tennis data
	aggregator := scrapers.NewAggregator(matchRepo, playerRepo, tournamentRepo)

	// 6a. Start Live Web Scraping Scheduler (runs every minute)
	scraperInterval := 1 * time.Minute
	if intervalStr := os.Getenv("SCRAPER_INTERVAL"); intervalStr != "" {
		if duration, err := time.ParseDuration(intervalStr); err == nil {
			scraperInterval = duration
		}
	}

	scraperScheduler := scraper.NewScheduler(tournamentRepo, playerRepo, matchRepo, scraperInterval)
	if err := scraperScheduler.Start(); err != nil {
		log.Printf("Warning: Failed to start scraper scheduler: %v", err)
	} else {
		log.Printf("✅ Scraper scheduler started (interval: %s)", scraperInterval)
	}

	// Check if simulator mode is enabled via environment variable
	enableSimulator := os.Getenv("ENABLE_SIMULATOR")

	// Cleanup: Delete simulated matches if simulator is disabled
	if enableSimulator != "true" {
		log.Println("ENABLE_SIMULATOR=false, cleaning up simulated matches from database")
		if err := matchRepo.DeleteSimulated(ctx); err != nil {
			log.Printf("Warning: Failed to delete simulated matches: %v", err)
		} else {
			log.Println("Successfully deleted all simulated matches")
		}
	}

	// Try to fetch real live matches first
	liveMatches, err := aggregator.FetchLiveMatches(ctx)
	if err != nil || len(liveMatches) == 0 {
		if enableSimulator == "true" {
			log.Printf("No real live matches available, ENABLE_SIMULATOR=true, starting simulator")

			// Fallback: Use simulator for demo/testing purposes
			sim := simulator.NewEngine(rdb, matchUpdateChan, matchRepo, playerRepo, tournamentRepo)
			sim.InitializeMatches()
			go sim.Start(context.Background())
		} else {
			log.Printf("No real live matches available. App will show empty state.")
			log.Printf("Set ENABLE_SIMULATOR=true to use simulator for testing.")
			// No matches - app will show "no live matches" message
		}
	} else {
		log.Printf("Found %d real live matches, starting periodic scraper", len(liveMatches))

		// Use real data: Start periodic fetching (every 30 seconds)
		go aggregator.StartPeriodicFetch(context.Background(), matchUpdateChan, 30*time.Second)

		// Send initial matches to WebSocket
		for _, match := range liveMatches {
			matchUpdateChan <- match
		}
	}

	// 7. Start Background Processes
	go hub.Run()

	// Bridge Simulator -> Websocket
	go func() {
		for match := range matchUpdateChan {
			hub.BroadcastMatchUpdate(match)
		}
	}()

	// 8. Handlers
	matchHandler := handlers.NewMatchHandler(matchRepo)
	tournamentHandler := handlers.NewTournamentHandler()

	// 9. Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// CORS Configuration for Railway deployment
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"*", // Allow all origins for development
			"https://hardcourt-production.up.railway.app",
			"https://accomplished-hope-production.up.railway.app",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health Check with Database connectivity
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		health := map[string]interface{}{
			"status": "healthy",
			"timestamp": time.Now(),
		}

		// Check database health
		if err := db.Health(r.Context()); err != nil {
			health["status"] = "unhealthy"
			health["database"] = "disconnected"
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			health["database"] = "connected"
		}

		// Check Redis health
		if _, err := rdb.Ping(r.Context()).Result(); err != nil {
			health["redis"] = "disconnected"
		} else {
			health["redis"] = "connected"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	})

	// REST API Routes
	r.Route("/api", func(r chi.Router) {
		// Match routes
		r.Get("/matches", matchHandler.GetAllMatches)
		r.Get("/matches/{id}", matchHandler.GetMatchByID)
		r.Get("/matches/{id}/highlights", tournamentHandler.GetMatchHighlights)
		r.Get("/matches/past", tournamentHandler.GetPastMatches)

		// Tournament routes
		r.Get("/tournaments", tournamentHandler.ListTournaments)
		r.Get("/tournaments/{id}", tournamentHandler.GetTournament)
		r.Get("/tournaments/{id}/matches", tournamentHandler.GetTournamentMatches)
		r.Get("/tournaments/{id}/draw", tournamentHandler.GetTournamentDraw)

		// Scraper monitoring endpoint
		r.Get("/scraper/status", func(w http.ResponseWriter, r *http.Request) {
			status := scraperScheduler.GetStatus()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(status)
		})
	})

	// WebSocket Route
	r.Get("/ws", hub.ServeWs)

	// 10. Server with Graceful Shutdown
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down gracefully...")

	// Stop scraper scheduler
	scraperScheduler.Stop()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped")
}
