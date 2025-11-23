package scrapers

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
	"golang.org/x/time/rate"
)

// Aggregator combines multiple data sources with fallback logic
type Aggregator struct {
	sofascore      *SofascoreClient
	matchRepo      *repository.MatchRepository
	playerRepo     *repository.PlayerRepository
	tournamentRepo *repository.TournamentRepository

	// Rate limiting
	limiter *rate.Limiter

	// Caching
	cache      map[string]*domain.Match
	cacheMu    sync.RWMutex
	cacheExpiry time.Duration
}

func NewAggregator(
	matchRepo *repository.MatchRepository,
	playerRepo *repository.PlayerRepository,
	tournamentRepo *repository.TournamentRepository,
) *Aggregator {
	return &Aggregator{
		sofascore:      NewSofascoreClient(),
		matchRepo:      matchRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
		limiter:        rate.NewLimiter(rate.Every(2*time.Second), 1), // 1 request every 2 seconds
		cache:          make(map[string]*domain.Match),
		cacheExpiry:    30 * time.Second,
	}
}

// FetchLiveMatches retrieves live matches from all available sources
func (a *Aggregator) FetchLiveMatches(ctx context.Context) ([]*domain.Match, error) {
	// Rate limiting
	if err := a.limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limit exceeded: %w", err)
	}

	// Try primary source: Sofascore
	matches, err := a.sofascore.GetLiveMatches()
	if err != nil {
		log.Printf("Sofascore fetch failed: %v, trying fallback...", err)

		// Fallback: Get from database
		matches, err = a.matchRepo.GetAll(ctx, string(domain.StatusLive))
		if err != nil {
			return nil, fmt.Errorf("all sources failed: %w", err)
		}

		log.Printf("Using %d cached matches from database", len(matches))
		return matches, nil
	}

	log.Printf("Fetched %d live matches from Sofascore", len(matches))

	// Persist to database
	for _, match := range matches {
		// Update cache
		a.cacheMu.Lock()
		a.cache[match.ID] = match
		a.cacheMu.Unlock()

		// Save to database
		if err := a.persistMatch(ctx, match); err != nil {
			log.Printf("Failed to persist match %s: %v", match.ID, err)
		}
	}

	return matches, nil
}

// persistMatch saves match data to the database
func (a *Aggregator) persistMatch(ctx context.Context, match *domain.Match) error {
	// Create/update tournament
	tournament := &domain.Tournament{
		ID:      match.TournamentID,
		Name:    match.TournamentID, // We'll enhance this later
		Surface: "Hard",              // Default, enhance later
		City:    "Unknown",
	}
	if err := a.tournamentRepo.Create(ctx, tournament); err != nil {
		return fmt.Errorf("failed to save tournament: %w", err)
	}

	// Create/update players
	if match.Player1 != nil {
		if err := a.playerRepo.Create(ctx, match.Player1); err != nil {
			return fmt.Errorf("failed to save player1: %w", err)
		}
	}
	if match.Player2 != nil {
		if err := a.playerRepo.Create(ctx, match.Player2); err != nil {
			return fmt.Errorf("failed to save player2: %w", err)
		}
	}

	// Check if match exists
	existing, err := a.matchRepo.GetByID(ctx, match.ID)
	if err != nil || existing == nil {
		// Create new match
		return a.matchRepo.Create(ctx, match)
	}

	// Update existing match
	return a.matchRepo.Update(ctx, match)
}

// StartPeriodicFetch runs continuous fetching in the background
func (a *Aggregator) StartPeriodicFetch(ctx context.Context, updateChan chan *domain.Match, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Printf("Starting periodic fetch every %v", interval)

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping periodic fetch")
			return
		case <-ticker.C:
			matches, err := a.FetchLiveMatches(ctx)
			if err != nil {
				log.Printf("Periodic fetch error: %v", err)
				continue
			}

			// Send updates via channel
			for _, match := range matches {
				select {
				case updateChan <- match:
				case <-ctx.Done():
					return
				default:
					// Channel full, skip this update
				}
			}
		}
	}
}

// GetCachedMatch retrieves a match from cache if available
func (a *Aggregator) GetCachedMatch(matchID string) (*domain.Match, bool) {
	a.cacheMu.RLock()
	defer a.cacheMu.RUnlock()

	match, exists := a.cache[matchID]
	return match, exists
}

// ClearExpiredCache removes old entries from cache
func (a *Aggregator) ClearExpiredCache() {
	a.cacheMu.Lock()
	defer a.cacheMu.Unlock()

	// Simple approach: clear entire cache periodically
	// In production, you'd track timestamps per entry
	a.cache = make(map[string]*domain.Match)
}
