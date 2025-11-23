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
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Surface    string     `json:"surface"` // Hard, Clay, Grass
	City       string     `json:"city"`
	Country    string     `json:"country"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	Year       int        `json:"year,omitempty"` // Year of tournament for filtering
	Category   string     `json:"category"`       // ATP/WTA, Grand Slam, Masters 1000
	PrizeMoney int64      `json:"prize_money"`
	Status     string     `json:"status"`          // upcoming, ongoing, completed
	WinnerID   *string    `json:"winner_id,omitempty"`    // Tournament champion
	RunnerUpID *string    `json:"runner_up_id,omitempty"` // Tournament finalist
	LogoURL    string     `json:"logo_url,omitempty"`     // Tournament logo
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// Player represents a tennis player
type Player struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
	Rank        int       `json:"rank"`
	Points      int       `json:"points"`
	Age         int       `json:"age,omitempty"`
	HeightCm    int       `json:"height_cm,omitempty"`
	WeightKg    int       `json:"weight_kg,omitempty"`
	Plays       string    `json:"plays,omitempty"`    // Right-handed, Left-handed
	Backhand    string    `json:"backhand,omitempty"` // One-handed, Two-handed
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Match represents a single match
type Match struct {
	ID              string      `json:"id"`
	TournamentID    string      `json:"tournament_id"`
	Tournament      *Tournament `json:"tournament,omitempty"`
	Player1ID       string      `json:"player1_id"`
	Player2ID       string      `json:"player2_id"`
	Player1         *Player     `json:"player1,omitempty"`
	Player2         *Player     `json:"player2,omitempty"`
	Status          MatchStatus `json:"status"`
	Round           string      `json:"round,omitempty"` // R128, R64, R32, R16, QF, SF, F
	StartTime       time.Time   `json:"start_time"`
	EndTime         *time.Time  `json:"end_time,omitempty"`
	WinnerID        *string     `json:"winner_id,omitempty"`
	DurationMinutes int         `json:"duration_minutes,omitempty"`
	Court           string      `json:"court,omitempty"`
	IsSimulated     bool        `json:"is_simulated"` // TRUE for simulator matches, FALSE for real
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`

	// Live Score Data
	Score ScoreState `json:"score"`
	Stats MatchStats `json:"stats"`
	Sets  []SetScore `json:"sets,omitempty"`

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
	AcesP1           int     `json:"aces_p1"`
	AcesP2           int     `json:"aces_p2"`
	DoubleFaultsP1   int     `json:"df_p1"`
	DoubleFaultsP2   int     `json:"df_p2"`
	BreakPointsP1    int     `json:"break_points_p1"`
	BreakPointsP2    int     `json:"break_points_p2"`
	WinnersP1        int     `json:"winners_p1"`
	WinnersP2        int     `json:"winners_p2"`
	UnforcedErrorsP1 int     `json:"unforced_errors_p1"`
	UnforcedErrorsP2 int     `json:"unforced_errors_p2"`
	FirstServePctP1  float64 `json:"first_serve_pct_p1"`
	FirstServePctP2  float64 `json:"first_serve_pct_p2"`
	RallyCount       int     `json:"rally_count"` // Current rally length simulation
}

// SetScore represents a completed set
type SetScore struct {
	SetNumber  int `json:"set_number"`
	GamesP1    int `json:"games_p1"`
	GamesP2    int `json:"games_p2"`
	TiebreakP1 int `json:"tiebreak_p1,omitempty"`
	TiebreakP2 int `json:"tiebreak_p2,omitempty"`
}

// TournamentDraw represents a position in the tournament bracket
type TournamentDraw struct {
	ID           int     `json:"id"`
	TournamentID string  `json:"tournament_id"`
	Round        string  `json:"round"`
	Position     int     `json:"position"`
	PlayerID     *string `json:"player_id,omitempty"`
	Player       *Player `json:"player,omitempty"`
	Seed         int     `json:"seed,omitempty"`
	Bye          bool    `json:"bye"`
}

// MatchHighlight represents a key moment in a match
type MatchHighlight struct {
	ID            int       `json:"id"`
	MatchID       string    `json:"match_id"`
	Timestamp     time.Time `json:"timestamp"`
	EventType     string    `json:"event_type"` // break_point, ace, rally, etc.
	Description   string    `json:"description"`
	LeverageIndex float64   `json:"leverage_index"`
}
