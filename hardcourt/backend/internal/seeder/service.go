package seeder

import (
	"context"
	"fmt"
	"log"
	"strings"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
)

// Service handles seeding of historical data
type Service struct {
	tournamentRepo *repository.TournamentRepository
	playerRepo     *repository.PlayerRepository
	playerCache    map[string]string // name -> player ID mapping
}

// NewService creates a new seeder service
func NewService(tournamentRepo *repository.TournamentRepository, playerRepo *repository.PlayerRepository) *Service {
	return &Service{
		tournamentRepo: tournamentRepo,
		playerRepo:     playerRepo,
		playerCache:    make(map[string]string),
	}
}

// SeedTournaments populates the database with historical tournament data
func (s *Service) SeedTournaments(ctx context.Context) error {
	log.Println("Starting tournament seeding...")

	data := GetTournamentSeedData()
	successCount := 0
	errorCount := 0

	for _, tournamentInfo := range data.Tournaments {
		if err := s.seedSingleTournament(ctx, tournamentInfo); err != nil {
			log.Printf("Warning: Failed to seed tournament %s (%s): %v",
				tournamentInfo.Name, tournamentInfo.Year, err)
			errorCount++
		} else {
			successCount++
		}
	}

	log.Printf("Tournament seeding complete: %d successful, %d errors", successCount, errorCount)

	if errorCount > 0 {
		return fmt.Errorf("seeding completed with %d errors", errorCount)
	}

	return nil
}

// seedSingleTournament seeds a single tournament with winner/runner-up data
func (s *Service) seedSingleTournament(ctx context.Context, info TournamentInfo) error {
	// Resolve winner and runner-up player IDs
	var winnerID, runnerUpID *string

	if info.WinnerName != "" {
		id, err := s.resolvePlayerID(ctx, info.WinnerName)
		if err != nil {
			return fmt.Errorf("failed to resolve winner %s: %w", info.WinnerName, err)
		}
		winnerID = &id
	}

	if info.RunnerUpName != "" {
		id, err := s.resolvePlayerID(ctx, info.RunnerUpName)
		if err != nil {
			return fmt.Errorf("failed to resolve runner-up %s: %w", info.RunnerUpName, err)
		}
		runnerUpID = &id
	}

	// Create tournament domain object
	tournament := &domain.Tournament{
		ID:         info.ID,
		Name:       info.Name,
		Surface:    info.Surface,
		City:       info.City,
		Country:    info.Country,
		StartDate:  &info.StartDate,
		EndDate:    &info.EndDate,
		Year:       info.Year,
		Category:   info.Category,
		PrizeMoney: info.PrizeMoney,
		Status:     info.Status,
		WinnerID:   winnerID,
		RunnerUpID: runnerUpID,
	}

	// Create or update tournament
	if err := s.tournamentRepo.Create(ctx, tournament); err != nil {
		return fmt.Errorf("failed to create tournament: %w", err)
	}

	log.Printf("✓ Seeded: %s %d (%s)", info.Name, info.Year, info.Category)
	return nil
}

// resolvePlayerID gets or creates a player ID from a player name
func (s *Service) resolvePlayerID(ctx context.Context, playerName string) (string, error) {
	// Check cache first
	if id, exists := s.playerCache[playerName]; exists {
		return id, nil
	}

	// Generate player ID from name (e.g., "J. Sinner" -> "j-sinner")
	playerID := s.generatePlayerID(playerName)

	// Check if player exists in database
	player, err := s.playerRepo.GetByID(ctx, playerID)
	if err == nil && player != nil {
		// Player exists, cache and return
		s.playerCache[playerName] = playerID
		return playerID, nil
	}

	// Player doesn't exist, create a basic entry
	newPlayer := &domain.Player{
		ID:          playerID,
		Name:        playerName,
		CountryCode: s.extractCountryCode(playerName),
		Rank:        0, // Will be updated by real data later
	}

	if err := s.playerRepo.Create(ctx, newPlayer); err != nil {
		// If error is duplicate, that's fine - another goroutine created it
		if !strings.Contains(err.Error(), "duplicate") && !strings.Contains(err.Error(), "conflict") {
			return "", fmt.Errorf("failed to create player: %w", err)
		}
	}

	// Cache and return
	s.playerCache[playerName] = playerID
	log.Printf("  → Created player: %s (ID: %s)", playerName, playerID)
	return playerID, nil
}

// generatePlayerID creates a URL-safe player ID from name
// Example: "J. Sinner" -> "j-sinner"
func (s *Service) generatePlayerID(name string) string {
	id := strings.ToLower(name)
	id = strings.ReplaceAll(id, ".", "")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.TrimSpace(id)
	return id
}

// extractCountryCode attempts to extract country code from player name
// This is a placeholder - in production you'd have a proper player database
func (s *Service) extractCountryCode(name string) string {
	// Map of known players to country codes
	countryMap := map[string]string{
		"J. Sinner":          "IT",
		"C. Alcaraz":         "ES",
		"N. Djokovic":        "RS",
		"D. Medvedev":        "RU",
		"A. Zverev":          "DE",
		"R. Nadal":           "ES",
		"S. Tsitsipas":       "GR",
		"D. Thiem":           "AT",
		"T. Fritz":           "US",
		"A. Rublev":          "RU",
		"H. Rune":            "DK",
		"C. Ruud":            "NO",
		"F. Auger-Aliassime": "CA",
		"N. Jarry":           "CL",
		"M. Arnaldi":         "IT",
		"G. Dimitrov":        "BG",
		"H. Hurkacz":         "PL",
		"J. Struff":          "DE",
		"A. de Minaur":       "AU",
		"N. Kyrgios":         "AU",
		"M. Berrettini":      "IT",
	}

	if code, exists := countryMap[name]; exists {
		return code
	}

	return "XX" // Unknown country
}

// SeedPlayers creates core ATP players with accurate data
func (s *Service) SeedPlayers(ctx context.Context) error {
	log.Println("Seeding core ATP players...")

	players := []domain.Player{
		{ID: "j-sinner", Name: "J. Sinner", CountryCode: "IT", Rank: 1, Points: 11180},
		{ID: "c-alcaraz", Name: "C. Alcaraz", CountryCode: "ES", Rank: 2, Points: 8500},
		{ID: "n-djokovic", Name: "N. Djokovic", CountryCode: "RS", Rank: 3, Points: 7900},
		{ID: "d-medvedev", Name: "D. Medvedev", CountryCode: "RU", Rank: 4, Points: 5000},
		{ID: "a-zverev", Name: "A. Zverev", CountryCode: "DE", Rank: 5, Points: 4800},
		{ID: "a-rublev", Name: "A. Rublev", CountryCode: "RU", Rank: 6, Points: 4100},
		{ID: "h-rune", Name: "H. Rune", CountryCode: "DK", Rank: 7, Points: 3800},
		{ID: "h-hurkacz", Name: "H. Hurkacz", CountryCode: "PL", Rank: 8, Points: 3500},
		{ID: "t-fritz", Name: "T. Fritz", CountryCode: "US", Rank: 9, Points: 3200},
		{ID: "s-tsitsipas", Name: "S. Tsitsipas", CountryCode: "GR", Rank: 10, Points: 3100},
		{ID: "c-ruud", Name: "C. Ruud", CountryCode: "NO", Rank: 11, Points: 3000},
		{ID: "g-dimitrov", Name: "G. Dimitrov", CountryCode: "BG", Rank: 12, Points: 2900},
		{ID: "r-nadal", Name: "R. Nadal", CountryCode: "ES", Rank: 150, Points: 500}, // Legend (injured/retired)
		{ID: "d-thiem", Name: "D. Thiem", CountryCode: "AT", Rank: 98, Points: 800},  // Former champion
	}

	successCount := 0
	for _, player := range players {
		if err := s.playerRepo.Create(ctx, &player); err != nil {
			log.Printf("Warning: Failed to seed player %s: %v", player.Name, err)
		} else {
			successCount++
		}
	}

	log.Printf("Player seeding complete: %d/%d successful", successCount, len(players))
	return nil
}
