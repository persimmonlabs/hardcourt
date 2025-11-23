package repository

import (
	"context"
	"fmt"

	"hardcourt/backend/internal/database"
	"hardcourt/backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type MatchRepository struct {
	db *database.DB
}

func NewMatchRepository(db *database.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

// Create inserts a new match into the database
func (r *MatchRepository) Create(ctx context.Context, match *domain.Match) error {
	query := `
		INSERT INTO matches (
			id, tournament_id, player1_id, player2_id, status, start_time,
			sets_p1, sets_p2, games_p1, games_p2, points_p1, points_p2, serving,
			win_prob_p1, leverage_index, fatigue_p1, fatigue_p2
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		ON CONFLICT (id) DO NOTHING
	`

	_, err := r.db.Pool.Exec(ctx, query,
		match.ID, match.TournamentID, match.Player1ID, match.Player2ID,
		match.Status, match.StartTime,
		match.Score.SetsP1, match.Score.SetsP2,
		match.Score.GamesP1, match.Score.GamesP2,
		match.Score.PointsP1, match.Score.PointsP2,
		match.Score.Serving,
		match.WinProbP1, match.LeverageIndex,
		match.FatigueP1, match.FatigueP2,
	)

	if err != nil {
		return fmt.Errorf("failed to create match: %w", err)
	}

	// Create stats record
	statsQuery := `
		INSERT INTO match_stats (match_id, aces_p1, aces_p2, df_p1, df_p2, break_points_p1, break_points_p2, rally_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (match_id) DO NOTHING
	`

	_, err = r.db.Pool.Exec(ctx, statsQuery,
		match.ID,
		match.Stats.AcesP1, match.Stats.AcesP2,
		match.Stats.DoubleFaultsP1, match.Stats.DoubleFaultsP2,
		match.Stats.BreakPointsP1, match.Stats.BreakPointsP2,
		match.Stats.RallyCount,
	)

	return err
}

// Update updates an existing match
func (r *MatchRepository) Update(ctx context.Context, match *domain.Match) error {
	query := `
		UPDATE matches SET
			status = $2, winner_id = $3,
			sets_p1 = $4, sets_p2 = $5, games_p1 = $6, games_p2 = $7,
			points_p1 = $8, points_p2 = $9, serving = $10,
			win_prob_p1 = $11, leverage_index = $12,
			fatigue_p1 = $13, fatigue_p2 = $14,
			updated_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.Pool.Exec(ctx, query,
		match.ID, match.Status, match.WinnerID,
		match.Score.SetsP1, match.Score.SetsP2,
		match.Score.GamesP1, match.Score.GamesP2,
		match.Score.PointsP1, match.Score.PointsP2,
		match.Score.Serving,
		match.WinProbP1, match.LeverageIndex,
		match.FatigueP1, match.FatigueP2,
	)

	if err != nil {
		return fmt.Errorf("failed to update match: %w", err)
	}

	// Update stats
	statsQuery := `
		UPDATE match_stats SET
			aces_p1 = $2, aces_p2 = $3, df_p1 = $4, df_p2 = $5,
			break_points_p1 = $6, break_points_p2 = $7, rally_count = $8
		WHERE match_id = $1
	`

	_, err = r.db.Pool.Exec(ctx, statsQuery,
		match.ID,
		match.Stats.AcesP1, match.Stats.AcesP2,
		match.Stats.DoubleFaultsP1, match.Stats.DoubleFaultsP2,
		match.Stats.BreakPointsP1, match.Stats.BreakPointsP2,
		match.Stats.RallyCount,
	)

	return err
}

// GetByID retrieves a match by ID with player information
func (r *MatchRepository) GetByID(ctx context.Context, id string) (*domain.Match, error) {
	query := `
		SELECT
			m.id, m.tournament_id, m.player1_id, m.player2_id, m.status, m.start_time, m.winner_id,
			m.sets_p1, m.sets_p2, m.games_p1, m.games_p2, m.points_p1, m.points_p2, m.serving,
			m.win_prob_p1, m.leverage_index, m.fatigue_p1, m.fatigue_p2,
			p1.id, p1.name, p1.country_code, p1.rank,
			p2.id, p2.name, p2.country_code, p2.rank,
			COALESCE(s.aces_p1, 0), COALESCE(s.aces_p2, 0),
			COALESCE(s.df_p1, 0), COALESCE(s.df_p2, 0),
			COALESCE(s.break_points_p1, 0), COALESCE(s.break_points_p2, 0),
			COALESCE(s.rally_count, 0)
		FROM matches m
		JOIN players p1 ON m.player1_id = p1.id
		JOIN players p2 ON m.player2_id = p2.id
		LEFT JOIN match_stats s ON m.id = s.match_id
		WHERE m.id = $1
	`

	match := &domain.Match{
		Player1: &domain.Player{},
		Player2: &domain.Player{},
	}

	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&match.ID, &match.TournamentID, &match.Player1ID, &match.Player2ID,
		&match.Status, &match.StartTime, &match.WinnerID,
		&match.Score.SetsP1, &match.Score.SetsP2,
		&match.Score.GamesP1, &match.Score.GamesP2,
		&match.Score.PointsP1, &match.Score.PointsP2,
		&match.Score.Serving,
		&match.WinProbP1, &match.LeverageIndex,
		&match.FatigueP1, &match.FatigueP2,
		&match.Player1.ID, &match.Player1.Name, &match.Player1.CountryCode, &match.Player1.Rank,
		&match.Player2.ID, &match.Player2.Name, &match.Player2.CountryCode, &match.Player2.Rank,
		&match.Stats.AcesP1, &match.Stats.AcesP2,
		&match.Stats.DoubleFaultsP1, &match.Stats.DoubleFaultsP2,
		&match.Stats.BreakPointsP1, &match.Stats.BreakPointsP2,
		&match.Stats.RallyCount,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("match not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get match: %w", err)
	}

	return match, nil
}

// GetAll retrieves all matches with optional status filter
func (r *MatchRepository) GetAll(ctx context.Context, status string) ([]*domain.Match, error) {
	query := `
		SELECT
			m.id, m.tournament_id, m.player1_id, m.player2_id, m.status, m.start_time, m.winner_id,
			m.sets_p1, m.sets_p2, m.games_p1, m.games_p2, m.points_p1, m.points_p2, m.serving,
			m.win_prob_p1, m.leverage_index, m.fatigue_p1, m.fatigue_p2,
			p1.id, p1.name, p1.country_code, p1.rank,
			p2.id, p2.name, p2.country_code, p2.rank,
			COALESCE(s.aces_p1, 0), COALESCE(s.aces_p2, 0),
			COALESCE(s.df_p1, 0), COALESCE(s.df_p2, 0),
			COALESCE(s.break_points_p1, 0), COALESCE(s.break_points_p2, 0),
			COALESCE(s.rally_count, 0)
		FROM matches m
		JOIN players p1 ON m.player1_id = p1.id
		JOIN players p2 ON m.player2_id = p2.id
		LEFT JOIN match_stats s ON m.id = s.match_id
	`

	var rows pgx.Rows
	var err error

	if status != "" {
		query += " WHERE m.status = $1 ORDER BY m.start_time DESC"
		rows, err = r.db.Pool.Query(ctx, query, status)
	} else {
		query += " ORDER BY m.start_time DESC"
		rows, err = r.db.Pool.Query(ctx, query)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query matches: %w", err)
	}
	defer rows.Close()

	var matches []*domain.Match
	for rows.Next() {
		match := &domain.Match{
			Player1: &domain.Player{},
			Player2: &domain.Player{},
		}

		err := rows.Scan(
			&match.ID, &match.TournamentID, &match.Player1ID, &match.Player2ID,
			&match.Status, &match.StartTime, &match.WinnerID,
			&match.Score.SetsP1, &match.Score.SetsP2,
			&match.Score.GamesP1, &match.Score.GamesP2,
			&match.Score.PointsP1, &match.Score.PointsP2,
			&match.Score.Serving,
			&match.WinProbP1, &match.LeverageIndex,
			&match.FatigueP1, &match.FatigueP2,
			&match.Player1.ID, &match.Player1.Name, &match.Player1.CountryCode, &match.Player1.Rank,
			&match.Player2.ID, &match.Player2.Name, &match.Player2.CountryCode, &match.Player2.Rank,
			&match.Stats.AcesP1, &match.Stats.AcesP2,
			&match.Stats.DoubleFaultsP1, &match.Stats.DoubleFaultsP2,
			&match.Stats.BreakPointsP1, &match.Stats.BreakPointsP2,
			&match.Stats.RallyCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan match: %w", err)
		}

		matches = append(matches, match)
	}

	return matches, rows.Err()
}
