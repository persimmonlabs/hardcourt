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
	seedPlayers := flag.Bool("players", false, "Seed ATP players only")
	seedTournaments := flag.Bool("tournaments", false, "Seed tournaments only")
	seedMatches := flag.Bool("matches", false, "Seed match results only")
	seedDraws := flag.Bool("draws", false, "Seed tournament draws only")
	seedAll := flag.Bool("all", false, "Seed everything (players, tournaments, matches, draws)")
	comprehensive := flag.Bool("comprehensive", true, "Use comprehensive dataset (ATP 500, ATP 250, Top 50 players)")
	flag.Parse()

	// Default to seeding all if no specific flags
	if !*seedPlayers && !*seedTournaments && !*seedMatches && !*seedDraws && !*seedAll {
		*seedAll = true
	}

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

	log.Println("✓ Connected to database")

	// Run migrations first
	log.Println("Running database migrations...")
	if err := db.RunMigrations(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("✓ Migrations completed")

	// Initialize repositories
	tournamentRepo := repository.NewTournamentRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	drawRepo := repository.NewTournamentDrawRepository(db)

	// Create seeder service
	seederService := seeder.NewService(tournamentRepo, playerRepo, matchRepo, drawRepo)

	// Determine mode
	mode := "standard"
	if *comprehensive {
		mode = "comprehensive"
	}
	log.Printf("Seeding mode: %s\n", mode)
	log.Println("===================================")

	// Execute seeding based on flags
	if *seedAll || *seedPlayers {
		log.Println("\n=== Seeding Players ===")
		if err := seederService.SeedPlayers(ctx, *comprehensive); err != nil {
			log.Printf("⚠️  Player seeding completed with errors: %v", err)
		} else {
			log.Println("✓ Player seeding successful!")
		}
	}

	if *seedAll || *seedTournaments {
		log.Println("\n=== Seeding Tournaments ===")
		if err := seederService.SeedTournaments(ctx, *comprehensive); err != nil {
			log.Printf("⚠️  Tournament seeding completed with errors: %v", err)
		} else {
			log.Println("✓ Tournament seeding successful!")
		}
	}

	if *seedAll || *seedMatches {
		log.Println("\n=== Seeding Match Results ===")
		if err := seederService.SeedMatches(ctx); err != nil {
			log.Printf("⚠️  Match seeding completed with errors: %v", err)
		} else {
			log.Println("✓ Match seeding successful!")
		}
	}

	if *seedAll || *seedDraws {
		log.Println("\n=== Seeding Tournament Draws ===")
		if err := seederService.SeedDraws(ctx); err != nil {
			log.Printf("⚠️  Draw seeding completed with errors: %v", err)
		} else {
			log.Println("✓ Draw seeding successful!")
		}
	}

	log.Println("\n===================================")
	log.Println("✓ Seeding complete!")
	log.Println("")

	// Summary
	if *comprehensive {
		log.Println("Comprehensive seeding completed:")
		log.Println("  • 70+ tournaments (Grand Slams, Masters 1000, ATP 500, ATP 250)")
		log.Println("  • 50+ players (Top 50 ATP + legends)")
		log.Println("  • 20+ Grand Slam finals (2020-2024)")
		log.Println("  • Tournament draws (seeded)")
	} else {
		log.Println("Standard seeding completed:")
		log.Println("  • 38 tournaments (Grand Slams, Masters 1000)")
		log.Println("  • 14 core players")
		log.Println("  • 20 Grand Slam finals (2020-2024)")
	}
}
