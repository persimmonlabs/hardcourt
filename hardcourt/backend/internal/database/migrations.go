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
		// Tournaments table with all fields
		`CREATE TABLE IF NOT EXISTS tournaments (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			surface VARCHAR(50) NOT NULL,
			city VARCHAR(255) NOT NULL,
			country VARCHAR(100),
			start_date DATE,
			end_date DATE,
			year INT,
			category VARCHAR(50),
			prize_money BIGINT,
			status VARCHAR(50) DEFAULT 'upcoming',
			winner_id VARCHAR(255),
			runner_up_id VARCHAR(255),
			logo_url VARCHAR(500),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)`,

		// Players table with full profile
		`CREATE TABLE IF NOT EXISTS players (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			country_code VARCHAR(3) NOT NULL,
			rank INT NOT NULL,
			points INT DEFAULT 0,
			age INT,
			height_cm INT,
			weight_kg INT,
			plays VARCHAR(20),
			backhand VARCHAR(20),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)`,

		// Matches table with is_simulated and all fields
		`CREATE TABLE IF NOT EXISTS matches (
			id VARCHAR(255) PRIMARY KEY,
			tournament_id VARCHAR(255) REFERENCES tournaments(id),
			player1_id VARCHAR(255) REFERENCES players(id),
			player2_id VARCHAR(255) REFERENCES players(id),
			status VARCHAR(50) NOT NULL,
			round VARCHAR(50),
			start_time TIMESTAMP WITH TIME ZONE NOT NULL,
			end_time TIMESTAMP WITH TIME ZONE,
			winner_id VARCHAR(255) REFERENCES players(id),
			duration_minutes INT,
			court VARCHAR(100),
			is_simulated BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

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
			fatigue_p2 DOUBLE PRECISION DEFAULT 0.0
		)`,

		// Match stats
		`CREATE TABLE IF NOT EXISTS match_stats (
			match_id VARCHAR(255) PRIMARY KEY REFERENCES matches(id) ON DELETE CASCADE,
			aces_p1 INT DEFAULT 0,
			aces_p2 INT DEFAULT 0,
			df_p1 INT DEFAULT 0,
			df_p2 INT DEFAULT 0,
			break_points_p1 INT DEFAULT 0,
			break_points_p2 INT DEFAULT 0,
			winners_p1 INT DEFAULT 0,
			winners_p2 INT DEFAULT 0,
			unforced_errors_p1 INT DEFAULT 0,
			unforced_errors_p2 INT DEFAULT 0,
			first_serve_pct_p1 FLOAT DEFAULT 0,
			first_serve_pct_p2 FLOAT DEFAULT 0,
			rally_count INT DEFAULT 0,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)`,

		// Tournament draws table
		`CREATE TABLE IF NOT EXISTS tournament_draws (
			id SERIAL PRIMARY KEY,
			tournament_id VARCHAR(255) REFERENCES tournaments(id),
			round VARCHAR(50) NOT NULL,
			position INT NOT NULL,
			player_id VARCHAR(255) REFERENCES players(id),
			seed INT,
			bye BOOLEAN DEFAULT false,
			UNIQUE(tournament_id, round, position)
		)`,

		// Match sets for historical data
		`CREATE TABLE IF NOT EXISTS match_sets (
			id SERIAL PRIMARY KEY,
			match_id VARCHAR(255) REFERENCES matches(id),
			set_number INT NOT NULL,
			games_p1 INT NOT NULL,
			games_p2 INT NOT NULL,
			tiebreak_p1 INT,
			tiebreak_p2 INT,
			UNIQUE(match_id, set_number)
		)`,

		// Match highlights
		`CREATE TABLE IF NOT EXISTS match_highlights (
			id SERIAL PRIMARY KEY,
			match_id VARCHAR(255) REFERENCES matches(id),
			timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			event_type VARCHAR(50),
			description TEXT,
			leverage_index FLOAT
		)`,

		// Add missing columns to existing tables (safe with IF NOT EXISTS)
		// Tournaments - add all potentially missing columns
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS year INT`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS category VARCHAR(50)`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS prize_money BIGINT`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'upcoming'`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS winner_id VARCHAR(255)`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS runner_up_id VARCHAR(255)`,
		`ALTER TABLE tournaments ADD COLUMN IF NOT EXISTS logo_url VARCHAR(500)`,

		// Players - add all potentially missing columns
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS points INT DEFAULT 0`,
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS age INT`,
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS height_cm INT`,
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS weight_kg INT`,
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS plays VARCHAR(20)`,
		`ALTER TABLE players ADD COLUMN IF NOT EXISTS backhand VARCHAR(20)`,

		// Matches - add all potentially missing columns
		`ALTER TABLE matches ADD COLUMN IF NOT EXISTS round VARCHAR(50)`,
		`ALTER TABLE matches ADD COLUMN IF NOT EXISTS duration_minutes INT`,
		`ALTER TABLE matches ADD COLUMN IF NOT EXISTS court VARCHAR(100)`,
		`ALTER TABLE matches ADD COLUMN IF NOT EXISTS is_simulated BOOLEAN DEFAULT FALSE`,

		// Indexes for performance
		`CREATE INDEX IF NOT EXISTS idx_tournaments_year ON tournaments(year DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_tournaments_status_year ON tournaments(status, year DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_tournaments_category ON tournaments(category)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_tournament ON matches(tournament_id)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_start_time ON matches(start_time DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_matches_simulated ON matches(is_simulated) WHERE is_simulated = TRUE`,
		`CREATE INDEX IF NOT EXISTS idx_highlights_match ON match_highlights(match_id)`,
	}

	for i, migration := range migrations {
		if _, err := db.Pool.Exec(ctx, migration); err != nil {
			return fmt.Errorf("migration %d failed: %w", i+1, err)
		}
	}

	log.Println("Migrations completed successfully")
	return nil
}
