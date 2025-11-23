package domain

import (
	"time"
)

// Enums for Match Status
type MatchStatus string

const (
	StatusScheduled MatchStatus = "Scheduled"
	StatusLive      MatchStatus = "Live"
	StatusFinished  MatchStatus = "Finished"
)

// Tournament represents a tennis tournament
type Tournament struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surface string `json:"surface"` // Hard, Clay, Grass
	City    string `json:"city"`
}

// Player represents a tennis player
type Player struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Rank        int    `json:"rank"`
}

// Match represents a single match
type Match struct {
	ID           string      `json:"id"`
	TournamentID string      `json:"tournament_id"`
	Player1ID    string      `json:"player1_id"`
	Player2ID    string      `json:"player2_id"`
	Player1      *Player     `json:"player1,omitempty"`
	Player2      *Player     `json:"player2,omitempty"`
	Status       MatchStatus `json:"status"`
	StartTime    time.Time   `json:"start_time"`
	WinnerID     *string     `json:"winner_id,omitempty"`

	// Live Score Data
	Score ScoreState `json:"score"`
	Stats MatchStats `json:"stats"`

	// Advanced Metrics
	WinProbP1     float64 `json:"win_prob_p1"`    // 0.0 to 1.0
	LeverageIndex float64 `json:"leverage_index"` // 0.0 to 1.0+
	FatigueP1     float64 `json:"fatigue_p1"`     // 0.0 to 100.0
	FatigueP2     float64 `json:"fatigue_p2"`     // 0.0 to 100.0
}

// ScoreState holds the current score
type ScoreState struct {
	SetsP1   int    `json:"sets_p1"`
	SetsP2   int    `json:"sets_p2"`
	GamesP1  int    `json:"games_p1"`
	GamesP2  int    `json:"games_p2"`
	PointsP1 string `json:"points_p1"` // "0", "15", "30", "40", "AD"
	PointsP2 string `json:"points_p2"`
	Serving  int    `json:"serving"` // 1 or 2
}

// MatchStats holds aggregate stats
type MatchStats struct {
	AcesP1         int `json:"aces_p1"`
	AcesP2         int `json:"aces_p2"`
	DoubleFaultsP1 int `json:"df_p1"`
	DoubleFaultsP2 int `json:"df_p2"`
	BreakPointsP1  int `json:"break_points_p1"`
	BreakPointsP2  int `json:"break_points_p2"`
	RallyCount     int `json:"rally_count"` // Current rally length simulation
}
