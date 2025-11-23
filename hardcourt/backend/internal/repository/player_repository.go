package repository

import (
	"context"
	"fmt"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/domain"
)

type PlayerRepository struct {
	db *database.DB
}

func NewPlayerRepository(db *database.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Create inserts a new player
func (r *PlayerRepository) Create(ctx context.Context, player *domain.Player) error {
	query := `
		INSERT INTO players (id, name, country_code, rank)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			country_code = EXCLUDED.country_code,
			rank = EXCLUDED.rank
	`

	_, err := r.db.Pool.Exec(ctx, query, player.ID, player.Name, player.CountryCode, player.Rank)
	if err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}

	return nil
}

// GetByID retrieves a player by ID
func (r *PlayerRepository) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	query := `SELECT id, name, country_code, rank FROM players WHERE id = $1`

	player := &domain.Player{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&player.ID, &player.Name, &player.CountryCode, &player.Rank,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get player: %w", err)
	}

	return player, nil
}
