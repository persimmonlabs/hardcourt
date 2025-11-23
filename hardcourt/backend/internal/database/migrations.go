package database

import (
	"context"
	"fmt"
	"log"
)

// RunMigrations executes the database schema migrations
func (db *DB) RunMigrations(ctx context.Context) error {
	log.Println("Running database migrations...")

	migrations := []string{
		`CREATE TABLE IF NOT EXISTS tournaments (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			surface VARCHAR(50) NOT NULL,
			city VARCHAR(255) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS players (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			country_code VARCHAR(3) NOT NULL,
			rank INT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS matches (
			id VARCHAR(255) PRIMARY KEY,
			tournament_id VARCHAR(255) REFERENCES tournaments(id),
			player1_id VARCHAR(255) REFERENCES players(id),
			player2_id VARCHAR(255) REFERENCES players(id),
			status VARCHAR(50) NOT NULL,
			start_time TIMESTAMP WITH TIME ZONE NOT NULL,
			winner_id VARCHAR(255) REFERENCES players(id),

			-- Score state
			sets_p1 INT DEFAULT 0,
			sets_p2 INT DEFAULT 0,
			games_p1 INT DEFAULT 0,
			games_p2 INT DEFAULT 0,
			points_p1 VARCHAR(10) DEFAULT '0',
			points_p2 VARCHAR(10) DEFAULT '0',
			serving INT DEFAULT 1,

			-- Advanced metrics
			win_prob_p1 DOUBLE PRECISION DEFAULT 0.5,
			leverage_index DOUBLE PRECISION DEFAULT 0.0,
			fatigue_p1 DOUBLE PRECISION DEFAULT 0.0,
			fatigue_p2 DOUBLE PRECISION DEFAULT 0.0,

			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS match_stats (
			match_id VARCHAR(255) PRIMARY KEY REFERENCES matches(id) ON DELETE CASCADE,
			aces_p1 INT DEFAULT 0,
			aces_p2 INT DEFAULT 0,
			df_p1 INT DEFAULT 0,
			df_p2 INT DEFAULT 0,
			break_points_p1 INT DEFAULT 0,
			break_points_p2 INT DEFAULT 0,
			rally_count INT DEFAULT 0
		)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_tournament ON matches(tournament_id)`,
	}

	for i, migration := range migrations {
		if _, err := db.Pool.Exec(ctx, migration); err != nil {
			return fmt.Errorf("migration %d failed: %w", i+1, err)
		}
	}

	log.Println("Migrations completed successfully")
	return nil
}
