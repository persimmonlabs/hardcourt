package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TournamentHandler struct {
	// Add repository dependency here later
}

func NewTournamentHandler() *TournamentHandler {
	return &TournamentHandler{}
}

// GET /api/tournaments - List all tournaments
func (h *TournamentHandler) ListTournaments(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status") // ongoing, completed, upcoming

	// TODO: Fetch from database
	// For now, return mock data
	tournaments := []map[string]interface{}{
		{
			"id":          "aus-open-2025",
			"name":        "Australian Open",
			"surface":     "Hard",
			"city":        "Melbourne",
			"country":     "Australia",
			"category":    "Grand Slam",
			"status":      "ongoing",
			"start_date":  "2025-01-14",
			"end_date":    "2025-01-27",
			"prize_money": 76500000,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tournaments)
}

// GET /api/tournaments/{id} - Get tournament details
func (h *TournamentHandler) GetTournament(w http.ResponseWriter, r *http.Request) {
	tournamentID := chi.URLParam(r, "id")

	// TODO: Fetch from database
	tournament := map[string]interface{}{
		"id":          tournamentID,
		"name":        "Australian Open",
		"surface":     "Hard",
		"city":        "Melbourne",
		"country":     "Australia",
		"category":    "Grand Slam",
		"status":      "ongoing",
		"start_date":  "2025-01-14",
		"end_date":    "2025-01-27",
		"prize_money": 76500000,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tournament)
}

// GET /api/tournaments/{id}/matches - Get all matches in tournament
func (h *TournamentHandler) GetTournamentMatches(w http.ResponseWriter, r *http.Request) {
	tournamentID := chi.URLParam(r, "id")
	round := r.URL.Query().Get("round") // R128, R64, R32, R16, QF, SF, F

	// TODO: Fetch from database filtered by tournament_id and optionally round
	matches := []map[string]interface{}{} // Empty for now

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tournament_id": tournamentID,
		"round":         round,
		"matches":       matches,
	})
}

// GET /api/tournaments/{id}/draw - Get tournament draw/bracket
func (h *TournamentHandler) GetTournamentDraw(w http.ResponseWriter, r *http.Request) {
	tournamentID := chi.URLParam(r, "id")

	// TODO: Fetch from tournament_draws table
	draw := map[string]interface{}{
		"tournament_id": tournamentID,
		"rounds": map[string][]interface{}{
			"R128": []interface{}{},
			"R64":  []interface{}{},
			"R32":  []interface{}{},
			"R16":  []interface{}{},
			"QF":   []interface{}{},
			"SF":   []interface{}{},
			"F":    []interface{}{},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(draw)
}

// GET /api/matches/past - Get historical matches
func (h *TournamentHandler) GetPastMatches(w http.ResponseWriter, r *http.Request) {
	// Query params: player, tournament, date_from, date_to, limit, offset
	playerID := r.URL.Query().Get("player")
	tournamentID := r.URL.Query().Get("tournament")

	// TODO: Fetch from database with filters
	matches := []map[string]interface{}{} // Empty for now

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"matches":     matches,
		"total":       0,
		"player_id":   playerID,
		"tournament_id": tournamentID,
	})
}

// GET /api/matches/{id}/highlights - Get match highlights
func (h *TournamentHandler) GetMatchHighlights(w http.ResponseWriter, r *http.Request) {
	matchID := chi.URLParam(r, "id")

	// TODO: Fetch from match_highlights table
	highlights := []map[string]interface{}{} // Empty for now

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"match_id":   matchID,
		"highlights": highlights,
	})
}
