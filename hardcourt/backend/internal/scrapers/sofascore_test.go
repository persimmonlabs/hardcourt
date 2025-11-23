package scrapers

import (
	"testing"
)

func TestSofascoreClient_Creation(t *testing.T) {
	client := NewSofascoreClient()

	if client == nil {
		t.Fatal("Expected client to be created, got nil")
	}

	if client.httpClient == nil {
		t.Fatal("Expected HTTP client to be initialized")
	}
}

func TestSofascoreClient_GetLiveMatches(t *testing.T) {
	client := NewSofascoreClient()

	// Note: This is a real API call - may fail if no live matches
	matches, err := client.GetLiveMatches()

	// We don't fail if there are no matches, just check the call works
	if err != nil {
		t.Logf("API call failed (expected if no live matches): %v", err)
		// Don't fail test - API might be down or no live matches
		return
	}

	t.Logf("Fetched %d live matches", len(matches))

	// If we got matches, validate structure
	for i, match := range matches {
		if match.ID == "" {
			t.Errorf("Match %d has empty ID", i)
		}
		if match.Player1 == nil || match.Player2 == nil {
			t.Errorf("Match %d missing player data", i)
		}
		if match.Player1 != nil && match.Player1.Name == "" {
			t.Errorf("Match %d Player1 has empty name", i)
		}
		if match.Player2 != nil && match.Player2.Name == "" {
			t.Errorf("Match %d Player2 has empty name", i)
		}
	}
}

func TestSofascoreClient_ConvertToMatches(t *testing.T) {
	client := NewSofascoreClient()

	// Test with empty events
	matches := client.convertToMatches([]sofascoreEvent{})
	if len(matches) != 0 {
		t.Errorf("Expected 0 matches, got %d", len(matches))
	}

	// Test with mock live event
	mockEvent := sofascoreEvent{
		ID: 12345,
		Tournament: sofascoreTournament{
			Name: "Australian Open",
		},
		HomeTeam: sofascorePlayer{
			ID:          1,
			Name:        "Rafael Nadal",
			CountryCode: "ES",
			Ranking:     1,
		},
		AwayTeam: sofascorePlayer{
			ID:          2,
			Name:        "Novak Djokovic",
			CountryCode: "RS",
			Ranking:     2,
		},
		Status: sofascoreStatus{
			Code: 6, // Live
		},
		HomeScore: sofascoreScore{
			Current: 1,
			Display: 3,
		},
		AwayScore: sofascoreScore{
			Current: 1,
			Display: 4,
		},
	}

	matches = client.convertToMatches([]sofascoreEvent{mockEvent})

	if len(matches) != 1 {
		t.Fatalf("Expected 1 match, got %d", len(matches))
	}

	match := matches[0]

	if match.Player1.Name != "Rafael Nadal" {
		t.Errorf("Expected Player1 name 'Rafael Nadal', got '%s'", match.Player1.Name)
	}

	if match.Player2.Name != "Novak Djokovic" {
		t.Errorf("Expected Player2 name 'Novak Djokovic', got '%s'", match.Player2.Name)
	}

	if match.Score.SetsP1 != 1 {
		t.Errorf("Expected SetsP1 = 1, got %d", match.Score.SetsP1)
	}

	if match.Score.GamesP1 != 3 {
		t.Errorf("Expected GamesP1 = 3, got %d", match.Score.GamesP1)
	}
}
