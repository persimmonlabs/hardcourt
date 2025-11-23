package repository

import (
	"context"
	"fmt"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/domain"
)

type TournamentDrawRepository struct {
	db *database.DB
}

func NewTournamentDrawRepository(db *database.DB) *TournamentDrawRepository {
	return &TournamentDrawRepository{db: db}
}

// Create inserts a new tournament draw entry
func (r *TournamentDrawRepository) Create(ctx context.Context, draw *domain.TournamentDraw) error {
	query := `
		INSERT INTO tournament_draws (
			tournament_id, round, position, player_id, seed, bye
		) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (tournament_id, round, position) DO NOTHING
	`

	_, err := r.db.Pool.Exec(ctx, query,
		draw.TournamentID, draw.Round, draw.Position,
		draw.PlayerID, draw.Seed, draw.Bye,
	)

	if err != nil {
		return fmt.Errorf("failed to create tournament draw: %w", err)
	}

	return nil
}

// GetByTournament retrieves all draw entries for a tournament
func (r *TournamentDrawRepository) GetByTournament(ctx context.Context, tournamentID string) ([]*domain.TournamentDraw, error) {
	query := `
		SELECT
			id, tournament_id, round, position, player_id, seed, bye
		FROM tournament_draws
		WHERE tournament_id = $1
		ORDER BY round, position
	`

	rows, err := r.db.Pool.Query(ctx, query, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("failed to query tournament draws: %w", err)
	}
	defer rows.Close()

	var draws []*domain.TournamentDraw
	for rows.Next() {
		draw := &domain.TournamentDraw{}

		err := rows.Scan(
			&draw.ID, &draw.TournamentID, &draw.Round,
			&draw.Position, &draw.PlayerID, &draw.Seed, &draw.Bye,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tournament draw: %w", err)
		}

		draws = append(draws, draw)
	}

	return draws, rows.Err()
}

// GetByRound retrieves draw entries for a specific round
func (r *TournamentDrawRepository) GetByRound(ctx context.Context, tournamentID, round string) ([]*domain.TournamentDraw, error) {
	query := `
		SELECT
			id, tournament_id, round, position, player_id, seed, bye
		FROM tournament_draws
		WHERE tournament_id = $1 AND round = $2
		ORDER BY position
	`

	rows, err := r.db.Pool.Query(ctx, query, tournamentID, round)
	if err != nil {
		return nil, fmt.Errorf("failed to query tournament draws for round: %w", err)
	}
	defer rows.Close()

	var draws []*domain.TournamentDraw
	for rows.Next() {
		draw := &domain.TournamentDraw{}

		err := rows.Scan(
			&draw.ID, &draw.TournamentID, &draw.Round,
			&draw.Position, &draw.PlayerID, &draw.Seed, &draw.Bye,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tournament draw: %w", err)
		}

		draws = append(draws, draw)
	}

	return draws, rows.Err()
}
