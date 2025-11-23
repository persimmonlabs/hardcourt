package seeder

import (
	"context"
	"fmt"
	"log"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
)

// DrawSeedData represents a tournament draw entry
type DrawSeedData struct {
	TournamentID string
	Round        string
	Position     int
	PlayerName   string
	Seed         int
	Bye          bool
}

// SeedDraws populates tournament draws/brackets
func (s *Service) SeedDraws(ctx context.Context) error {
	log.Println("Starting draw seeding...")

	// Generate draw data for key tournaments
	drawData := s.generateTournamentDraws()

	successCount := 0
	errorCount := 0

	for _, draw := range drawData {
		if err := s.seedSingleDraw(ctx, draw); err != nil {
			log.Printf("Warning: Failed to seed draw: %v", err)
			errorCount++
		} else {
			successCount++
		}
	}

	log.Printf("Draw seeding complete: %d successful, %d errors", successCount, errorCount)
	return nil
}

func (s *Service) seedSingleDraw(ctx context.Context, drawData DrawSeedData) error {
	var playerID *string

	if drawData.PlayerName != "" {
		pid, err := s.resolvePlayerID(ctx, drawData.PlayerName)
		if err != nil {
			return fmt.Errorf("failed to resolve player %s: %w", drawData.PlayerName, err)
		}
		playerID = &pid
	}

	// Create draw entry
	draw := &domain.TournamentDraw{
		TournamentID: drawData.TournamentID,
		Round:        drawData.Round,
		Position:     drawData.Position,
		PlayerID:     playerID,
		Seed:         drawData.Seed,
		Bye:          drawData.Bye,
	}

	if err := s.drawRepo.Create(ctx, draw); err != nil {
		return fmt.Errorf("failed to create draw entry: %w", err)
	}

	return nil
}

// generateTournamentDraws creates draw data for 2024 Australian Open (example)
func (s *Service) generateTournamentDraws() []DrawSeedData {
	ausOpen2024 := []DrawSeedData{
		// Top Half - Seeded players
		{TournamentID: "aus-open-2024", Round: "R128", Position: 1, PlayerName: "N. Djokovic", Seed: 1},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 2, PlayerName: "J. Sinner", Seed: 4},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 3, PlayerName: "D. Medvedev", Seed: 3},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 4, PlayerName: "C. Alcaraz", Seed: 2},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 5, PlayerName: "A. Rublev", Seed: 5},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 6, PlayerName: "S. Tsitsipas", Seed: 7},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 7, PlayerName: "A. Zverev", Seed: 6},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 8, PlayerName: "H. Hurkacz", Seed: 8},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 9, PlayerName: "H. Rune", Seed: 8},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 10, PlayerName: "T. Fritz", Seed: 10},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 11, PlayerName: "G. Dimitrov", Seed: 11},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 12, PlayerName: "T. Paul", Seed: 12},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 13, PlayerName: "A. de Minaur", Seed: 13},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 14, PlayerName: "U. Humbert", Seed: 14},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 15, PlayerName: "K. Khachanov", Seed: 15},
		{TournamentID: "aus-open-2024", Round: "R128", Position: 16, PlayerName: "B. Shelton", Seed: 16},
	}

	// Finals progression
	finals := []DrawSeedData{
		{TournamentID: "aus-open-2024", Round: "SF", Position: 1, PlayerName: "J. Sinner", Seed: 4},
		{TournamentID: "aus-open-2024", Round: "SF", Position: 2, PlayerName: "N. Djokovic", Seed: 1},
		{TournamentID: "aus-open-2024", Round: "SF", Position: 3, PlayerName: "D. Medvedev", Seed: 3},
		{TournamentID: "aus-open-2024", Round: "SF", Position: 4, PlayerName: "A. Zverev", Seed: 6},

		{TournamentID: "aus-open-2024", Round: "F", Position: 1, PlayerName: "J. Sinner", Seed: 4},
		{TournamentID: "aus-open-2024", Round: "F", Position: 2, PlayerName: "D. Medvedev", Seed: 3},
	}

	return append(ausOpen2024, finals...)
}
