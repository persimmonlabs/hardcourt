package repository

import (
	"context"
	"fmt"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/domain"
)

type TournamentRepository struct {
	db *database.DB
}

func NewTournamentRepository(db *database.DB) *TournamentRepository {
	return &TournamentRepository{db: db}
}

// Create inserts a new tournament
func (r *TournamentRepository) Create(ctx context.Context, tournament *domain.Tournament) error {
	query := `
		INSERT INTO tournaments (id, name, surface, city)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			surface = EXCLUDED.surface,
			city = EXCLUDED.city
	`

	_, err := r.db.Pool.Exec(ctx, query, tournament.ID, tournament.Name, tournament.Surface, tournament.City)
	if err != nil {
		return fmt.Errorf("failed to create tournament: %w", err)
	}

	return nil
}

// GetByID retrieves a tournament by ID
func (r *TournamentRepository) GetByID(ctx context.Context, id string) (*domain.Tournament, error) {
	query := `SELECT id, name, surface, city FROM tournaments WHERE id = $1`

	tournament := &domain.Tournament{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&tournament.ID, &tournament.Name, &tournament.Surface, &tournament.City,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get tournament: %w", err)
	}

	return tournament, nil
}
