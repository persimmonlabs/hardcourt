package scrapers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"hardcourt/backend/internal/domain"
)

const (
	sofascoreBaseURL = "https://api.sofascore.com/api/v1"
	userAgent        = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
)

type SofascoreClient struct {
	httpClient *http.Client
}

func NewSofascoreClient() *SofascoreClient {
	return &SofascoreClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Sofascore API response structures
type sofascoreResponse struct {
	Events []sofascoreEvent `json:"events"`
}

type sofascoreEvent struct {
	ID               int                    `json:"id"`
	Tournament       sofascoreTournament    `json:"tournament"`
	HomeTeam         sofascorePlayer        `json:"homeTeam"`
	AwayTeam         sofascorePlayer        `json:"awayTeam"`
	Status           sofascoreStatus        `json:"status"`
	HomeScore        sofascoreScore         `json:"homeScore"`
	AwayScore        sofascoreScore         `json:"awayScore"`
	StartTimestamp   int64                  `json:"startTimestamp"`
	WinnerCode       int                    `json:"winnerCode,omitempty"`
}

type sofascoreTournament struct {
	Name    string `json:"name"`
	Surface string `json:"groundType,omitempty"`
	City    string `json:"city,omitempty"`
}

type sofascorePlayer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Ranking     int    `json:"ranking,omitempty"`
}

type sofascoreStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type sofascoreScore struct {
	Current int   `json:"current"` // Sets won
	Period1 int   `json:"period1,omitempty"`
	Period2 int   `json:"period2,omitempty"`
	Period3 int   `json:"period3,omitempty"`
	Period4 int   `json:"period4,omitempty"`
	Period5 int   `json:"period5,omitempty"`
	Display int   `json:"display,omitempty"` // Current game score
}

// GetLiveMatches fetches currently live tennis matches
func (s *SofascoreClient) GetLiveMatches() ([]*domain.Match, error) {
	url := fmt.Sprintf("%s/sport/tennis/scheduled-events/%s", sofascoreBaseURL, time.Now().Format("2006-01-02"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch live matches: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("sofascore API returned %d: %s", resp.StatusCode, string(body))
	}

	var apiResp sofascoreResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return s.convertToMatches(apiResp.Events), nil
}

// convertToMatches converts Sofascore events to our domain Match model
func (s *SofascoreClient) convertToMatches(events []sofascoreEvent) []*domain.Match {
	matches := make([]*domain.Match, 0, len(events))

	for _, event := range events {
		// Only include live matches (status code 6 is "inprogress")
		if event.Status.Code != 6 {
			continue
		}

		match := &domain.Match{
			ID:           fmt.Sprintf("sofa_%d", event.ID),
			TournamentID: fmt.Sprintf("t_%s", event.Tournament.Name),
			Player1ID:    fmt.Sprintf("p_%d", event.HomeTeam.ID),
			Player2ID:    fmt.Sprintf("p_%d", event.AwayTeam.ID),
			Status:       domain.StatusLive,
			StartTime:    time.Unix(event.StartTimestamp, 0),
			Player1: &domain.Player{
				ID:          fmt.Sprintf("p_%d", event.HomeTeam.ID),
				Name:        event.HomeTeam.Name,
				CountryCode: event.HomeTeam.CountryCode,
				Rank:        event.HomeTeam.Ranking,
			},
			Player2: &domain.Player{
				ID:          fmt.Sprintf("p_%d", event.AwayTeam.ID),
				Name:        event.AwayTeam.Name,
				CountryCode: event.AwayTeam.CountryCode,
				Rank:        event.AwayTeam.Ranking,
			},
			Score: domain.ScoreState{
				SetsP1:   event.HomeScore.Current,
				SetsP2:   event.AwayScore.Current,
				GamesP1:  event.HomeScore.Display,
				GamesP2:  event.AwayScore.Display,
				PointsP1: "0", // Sofascore doesn't provide point-by-point
				PointsP2: "0",
				Serving:  1, // Default
			},
			Stats:     domain.MatchStats{},
			WinProbP1: 0.5, // Calculate separately if needed
		}

		// Determine winner if match is finished
		if event.Status.Code == 7 { // Finished
			match.Status = domain.StatusFinished
			if event.WinnerCode == 1 {
				winnerID := match.Player1ID
				match.WinnerID = &winnerID
			} else if event.WinnerCode == 2 {
				winnerID := match.Player2ID
				match.WinnerID = &winnerID
			}
		}

		matches = append(matches, match)
	}

	return matches
}

// GetMatchDetails fetches detailed stats for a specific match
func (s *SofascoreClient) GetMatchDetails(matchID int) (*domain.Match, error) {
	url := fmt.Sprintf("%s/event/%d", sofascoreBaseURL, matchID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch match details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sofascore API returned %d", resp.StatusCode)
	}

	var apiResp struct {
		Event sofascoreEvent `json:"event"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	matches := s.convertToMatches([]sofascoreEvent{apiResp.Event})
	if len(matches) == 0 {
		return nil, fmt.Errorf("no match found")
	}

	return matches[0], nil
}
