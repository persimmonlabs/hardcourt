package main

import (
	"context"
	"flag"
	"log"
	"os"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/repository"
	"hardcourt/backend/internal/seeder"
)

func main() {
	// Parse command-line flags
	seedPlayers := flag.Bool("players", false, "Seed core ATP players")
	seedTournaments := flag.Bool("tournaments", false, "Seed historical tournaments")
	seedAll := flag.Bool("all", true, "Seed both players and tournaments (default)")
	flag.Parse()

	// Database connection
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"
	}

	ctx := context.Background()

	// Connect to database
	db, err := database.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to database")

	// Run migrations first
	if err := db.RunMigrations(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations complete")

	// Initialize repositories
	tournamentRepo := repository.NewTournamentRepository(db)
	playerRepo := repository.NewPlayerRepository(db)

	// Create seeder service
	seederService := seeder.NewService(tournamentRepo, playerRepo)

	// Execute seeding based on flags
	if *seedAll || *seedPlayers {
		log.Println("\n=== Seeding Players ===")
		if err := seederService.SeedPlayers(ctx); err != nil {
			log.Printf("Player seeding completed with errors: %v", err)
		} else {
			log.Println("Player seeding successful!")
		}
	}

	if *seedAll || *seedTournaments {
		log.Println("\n=== Seeding Tournaments ===")
		if err := seederService.SeedTournaments(ctx); err != nil {
			log.Printf("Tournament seeding completed with errors: %v", err)
		} else {
			log.Println("Tournament seeding successful!")
		}
	}

	log.Println("\n✓ Seeding complete!")
}
