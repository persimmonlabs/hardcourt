package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
)

// MatchSeedData represents a single match result
type MatchSeedData struct {
	TournamentID string
	Round        string
	Player1Name  string
	Player2Name  string
	WinnerName   string
	Score        ScoreData
	Date         time.Time
	DurationMins int
}

type ScoreData struct {
	SetsP1   int
	SetsP2   int
	GamesP1  []int // Games per set
	GamesP2  []int
}

// SeedMatches populates matches for tournaments
func (s *Service) SeedMatches(ctx context.Context) error {
	log.Println("Starting match seeding...")

	// Example: Seed finals for each tournament
	matchData := s.generateTournamentMatches()

	successCount := 0
	errorCount := 0

	for _, match := range matchData {
		if err := s.seedSingleMatch(ctx, match); err != nil {
			log.Printf("Warning: Failed to seed match: %v", err)
			errorCount++
		} else {
			successCount++
		}
	}

	log.Printf("Match seeding complete: %d successful, %d errors", successCount, errorCount)
	return nil
}

func (s *Service) seedSingleMatch(ctx context.Context, matchData MatchSeedData) error {
	// Resolve player IDs
	player1ID, err := s.resolvePlayerID(ctx, matchData.Player1Name)
	if err != nil {
		return fmt.Errorf("failed to resolve player1 %s: %w", matchData.Player1Name, err)
	}

	player2ID, err := s.resolvePlayerID(ctx, matchData.Player2Name)
	if err != nil {
		return fmt.Errorf("failed to resolve player2 %s: %w", matchData.Player2Name, err)
	}

	winnerID, err := s.resolvePlayerID(ctx, matchData.WinnerName)
	if err != nil {
		return fmt.Errorf("failed to resolve winner %s: %w", matchData.WinnerName, err)
	}

	// Generate match ID
	matchID := fmt.Sprintf("%s-%s-vs-%s", matchData.TournamentID, player1ID, player2ID)

	// Create match
	match := &domain.Match{
		ID:              matchID,
		TournamentID:    matchData.TournamentID,
		Player1ID:       player1ID,
		Player2ID:       player2ID,
		Status:          domain.StatusFinished,
		Round:           matchData.Round,
		StartTime:       matchData.Date,
		WinnerID:        &winnerID,
		DurationMinutes: matchData.DurationMins,
		IsSimulated:     false,
		Score: domain.ScoreState{
			SetsP1:  matchData.Score.SetsP1,
			SetsP2:  matchData.Score.SetsP2,
			Serving: 0,
		},
	}

	if err := s.matchRepo.Create(ctx, match); err != nil {
		return fmt.Errorf("failed to create match: %w", err)
	}

	return nil
}

// generateTournamentMatches creates finals data for major tournaments
func (s *Service) generateTournamentMatches() []MatchSeedData {
	return []MatchSeedData{
		// 2024 Grand Slam Finals
		{
			TournamentID: "aus-open-2024", Round: "F",
			Player1Name: "J. Sinner", Player2Name: "D. Medvedev", WinnerName: "J. Sinner",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 6, 4, 6, 3}, GamesP2: []int{3, 3, 6, 4, 6}},
			Date: time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), DurationMins: 210,
		},
		{
			TournamentID: "roland-garros-2024", Round: "F",
			Player1Name: "C. Alcaraz", Player2Name: "A. Zverev", WinnerName: "C. Alcaraz",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 2, 6, 1, 6}, GamesP2: []int{3, 6, 5, 6, 2}},
			Date: time.Date(2024, 6, 9, 0, 0, 0, 0, time.UTC), DurationMins: 260,
		},
		{
			TournamentID: "wimbledon-2024", Round: "F",
			Player1Name: "C. Alcaraz", Player2Name: "N. Djokovic", WinnerName: "C. Alcaraz",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 2, 4}},
			Date: time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), DurationMins: 165,
		},
		{
			TournamentID: "us-open-2024", Round: "F",
			Player1Name: "J. Sinner", Player2Name: "T. Fritz", WinnerName: "J. Sinner",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 7}, GamesP2: []int{3, 4, 5}},
			Date: time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), DurationMins: 140,
		},

		// 2023 Grand Slam Finals
		{
			TournamentID: "aus-open-2023", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "S. Tsitsipas", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 7}, GamesP2: []int{3, 6, 6}},
			Date: time.Date(2023, 1, 29, 0, 0, 0, 0, time.UTC), DurationMins: 180,
		},
		{
			TournamentID: "roland-garros-2023", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "C. Ruud", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 6, 7}, GamesP2: []int{6, 3, 5}},
			Date: time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC), DurationMins: 195,
		},
		{
			TournamentID: "wimbledon-2023", Round: "F",
			Player1Name: "C. Alcaraz", Player2Name: "N. Djokovic", WinnerName: "C. Alcaraz",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{1, 7, 6, 3, 6}, GamesP2: []int{6, 6, 1, 6, 4}},
			Date: time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC), DurationMins: 288,
		},
		{
			TournamentID: "us-open-2023", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "D. Medvedev", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 6}, GamesP2: []int{3, 6, 3}},
			Date: time.Date(2023, 9, 10, 0, 0, 0, 0, time.UTC), DurationMins: 178,
		},

		// 2022 Grand Slam Finals
		{
			TournamentID: "aus-open-2022", Round: "F",
			Player1Name: "R. Nadal", Player2Name: "D. Medvedev", WinnerName: "R. Nadal",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{2, 6, 6, 6, 7}, GamesP2: []int{6, 7, 4, 4, 5}},
			Date: time.Date(2022, 1, 30, 0, 0, 0, 0, time.UTC), DurationMins: 330,
		},
		{
			TournamentID: "roland-garros-2022", Round: "F",
			Player1Name: "R. Nadal", Player2Name: "C. Ruud", WinnerName: "R. Nadal",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 3, 0}},
			Date: time.Date(2022, 6, 5, 0, 0, 0, 0, time.UTC), DurationMins: 135,
		},
		{
			TournamentID: "wimbledon-2022", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "N. Kyrgios", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{4, 6, 6, 7}, GamesP2: []int{6, 3, 4, 6}},
			Date: time.Date(2022, 7, 10, 0, 0, 0, 0, time.UTC), DurationMins: 186,
		},
		{
			TournamentID: "us-open-2022", Round: "F",
			Player1Name: "C. Alcaraz", Player2Name: "C. Ruud", WinnerName: "C. Alcaraz",
			Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 2, 7, 6}, GamesP2: []int{4, 6, 6, 3}},
			Date: time.Date(2022, 9, 11, 0, 0, 0, 0, time.UTC), DurationMins: 215,
		},

		// 2021 Grand Slam Finals
		{
			TournamentID: "aus-open-2021", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "D. Medvedev", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 6, 7}, GamesP2: []int{5, 2, 5}},
			Date: time.Date(2021, 2, 21, 0, 0, 0, 0, time.UTC), DurationMins: 113,
		},
		{
			TournamentID: "roland-garros-2021", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "S. Tsitsipas", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 2, 6, 6, 6}, GamesP2: []int{7, 6, 3, 2, 4}},
			Date: time.Date(2021, 6, 13, 0, 0, 0, 0, time.UTC), DurationMins: 255,
		},
		{
			TournamentID: "wimbledon-2021", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "M. Berrettini", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{7, 4, 4, 3}},
			Date: time.Date(2021, 7, 11, 0, 0, 0, 0, time.UTC), DurationMins: 205,
		},
		{
			TournamentID: "us-open-2021", Round: "F",
			Player1Name: "D. Medvedev", Player2Name: "N. Djokovic", WinnerName: "D. Medvedev",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{4, 4, 4}},
			Date: time.Date(2021, 9, 12, 0, 0, 0, 0, time.UTC), DurationMins: 135,
		},

		// 2020 Grand Slam Finals
		{
			TournamentID: "aus-open-2020", Round: "F",
			Player1Name: "N. Djokovic", Player2Name: "D. Thiem", WinnerName: "N. Djokovic",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 4, 2, 6, 6}, GamesP2: []int{4, 6, 6, 4, 4}},
			Date: time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC), DurationMins: 238,
		},
		{
			TournamentID: "roland-garros-2020", Round: "F",
			Player1Name: "R. Nadal", Player2Name: "N. Djokovic", WinnerName: "R. Nadal",
			Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 7}, GamesP2: []int{0, 2, 5}},
			Date: time.Date(2020, 10, 11, 0, 0, 0, 0, time.UTC), DurationMins: 159,
		},
		{
			TournamentID: "us-open-2020", Round: "F",
			Player1Name: "D. Thiem", Player2Name: "A. Zverev", WinnerName: "D. Thiem",
			Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{2, 4, 6, 6, 7}, GamesP2: []int{6, 6, 4, 3, 6}},
			Date: time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC), DurationMins: 251,
		},
	}
}
