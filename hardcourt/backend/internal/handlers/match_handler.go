package handlers

import (
	"encoding/json"
	"net/http"

	"hardcourt/backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

type MatchHandler struct {
	matchRepo *repository.MatchRepository
}

func NewMatchHandler(matchRepo *repository.MatchRepository) *MatchHandler {
	return &MatchHandler{matchRepo: matchRepo}
}

// GetAllMatches handles GET /api/matches
func (h *MatchHandler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	matches, err := h.matchRepo.GetAll(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

// GetMatchByID handles GET /api/matches/{id}
func (h *MatchHandler) GetMatchByID(w http.ResponseWriter, r *http.Request) {
	matchID := chi.URLParam(r, "id")

	match, err := h.matchRepo.GetByID(r.Context(), matchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}
