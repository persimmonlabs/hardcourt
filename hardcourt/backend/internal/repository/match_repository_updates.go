package repository

import (
	"context"
	"fmt"
)

// DeleteSimulated deletes all simulated matches from the database
func (r *MatchRepository) DeleteSimulated(ctx context.Context) error {
	query := `DELETE FROM matches WHERE is_simulated = TRUE`

	result, err := r.db.Pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete simulated matches: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected > 0 {
		fmt.Printf("Deleted %d simulated matches\n", rowsAffected)
	}

	return nil
}

// Note: The Create, Update, GetByID, and GetAll methods in match_repository.go
// need to be updated to include the is_simulated column in their SQL queries.
// Add is_simulated to INSERT and SELECT statements.
