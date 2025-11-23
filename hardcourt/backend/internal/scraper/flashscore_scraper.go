package scraper

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
)

// FlashScoreScraper scrapes live tennis scores from FlashScore
type FlashScoreScraper struct {
	matchRepo  *repository.MatchRepository
	playerRepo *repository.PlayerRepository
	httpClient *http.Client
}

// NewFlashScoreScraper creates a new FlashScore scraper
func NewFlashScoreScraper(
	matchRepo *repository.MatchRepository,
	playerRepo *repository.PlayerRepository,
) *FlashScoreScraper {
	return &FlashScoreScraper{
		matchRepo:  matchRepo,
		playerRepo: playerRepo,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
		},
	}
}

const (
	FlashScoreTennisURL = "https://www.flashscore.com/tennis/"
)

// ScrapeLiveMatches scrapes live tennis matches from FlashScore
func (s *FlashScoreScraper) ScrapeLiveMatches(ctx context.Context) error {
	log.Println("⚡ Scraping live matches from FlashScore...")

	req, err := http.NewRequestWithContext(ctx, "GET", FlashScoreTennisURL, nil)
	if err != nil {
		return err
	}

	// Mimic browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch FlashScore: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	updateCount := 0

	// Parse live matches
	doc.Find(".event__match, .sportName.tennis .event").Each(func(i int, match *goquery.Selection) {
		// Extract match data
		player1 := strings.TrimSpace(match.Find(".event__participant--home").Text())
		player2 := strings.TrimSpace(match.Find(".event__participant--away").Text())
		score1 := strings.TrimSpace(match.Find(".event__score--home").Text())
		score2 := strings.TrimSpace(match.Find(".event__score--away").Text())
		status := strings.TrimSpace(match.Find(".event__stage").Text())
		_ = strings.TrimSpace(match.Find(".event__title").Text()) // tournament - for future use

		if player1 == "" || player2 == "" {
			return
		}

		player1ID := s.generatePlayerID(player1)
		player2ID := s.generatePlayerID(player2)
		matchID := fmt.Sprintf("live-%s-vs-%s-%d", player1ID, player2ID, time.Now().Unix())

		// Determine match status
		matchStatus := domain.StatusLive
		if strings.Contains(strings.ToLower(status), "finished") ||
			strings.Contains(strings.ToLower(status), "ended") {
			matchStatus = domain.StatusFinished
		} else if strings.Contains(strings.ToLower(status), "not started") ||
			strings.Contains(strings.ToLower(status), "scheduled") {
			matchStatus = domain.StatusScheduled
		}

		matchObj := &domain.Match{
			ID:          matchID,
			Player1ID:   player1ID,
			Player2ID:   player2ID,
			Status:      matchStatus,
			StartTime:   time.Now(),
			IsSimulated: false,
		}

		// Try to parse scores
		if score1 != "" && score2 != "" {
			// Parse set scores (e.g., "6", "7", "6")
			// This is simplified - real implementation would parse detailed scores
		}

		// Check if match exists
		existing, err := s.matchRepo.GetByID(ctx, matchID)
		if err != nil || existing == nil {
			if err := s.matchRepo.Create(ctx, matchObj); err != nil {
				if !strings.Contains(err.Error(), "duplicate") && !strings.Contains(err.Error(), "foreign key") {
					log.Printf("Failed to create match from FlashScore: %v", err)
				}
			} else {
				updateCount++
			}
		}
	})

	log.Printf("✅ Scraped %d live matches from FlashScore", updateCount)
	return nil
}

func (s *FlashScoreScraper) generatePlayerID(name string) string {
	id := strings.ToLower(name)
	id = strings.ReplaceAll(id, ".", "")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.TrimSpace(id)
	return id
}
