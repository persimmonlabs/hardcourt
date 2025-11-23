package scraper

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/repository"
)

// ATPTourScraper scrapes live data from atptour.com
type ATPTourScraper struct {
	tournamentRepo *repository.TournamentRepository
	playerRepo     *repository.PlayerRepository
	matchRepo      *repository.MatchRepository
	httpClient     *http.Client
}

// NewATPTourScraper creates a new ATP Tour scraper
func NewATPTourScraper(
	tournamentRepo *repository.TournamentRepository,
	playerRepo *repository.PlayerRepository,
	matchRepo *repository.MatchRepository,
) *ATPTourScraper {
	return &ATPTourScraper{
		tournamentRepo: tournamentRepo,
		playerRepo:     playerRepo,
		matchRepo:      matchRepo,
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

// ATP Tour URLs
const (
	ATPRankingsPage    = "https://www.atptour.com/en/rankings/singles"
	ATPTournamentsPage = "https://www.atptour.com/en/tournaments"
	ATPScoresPage      = "https://www.atptour.com/en/scores/current"
)

// ScrapeAll scrapes all live ATP data
func (s *ATPTourScraper) ScrapeAll(ctx context.Context) error {
	log.Println("🕷️  Starting live ATP Tour web scraping...")

	// Scrape rankings
	if err := s.ScrapeRankings(ctx); err != nil {
		log.Printf("⚠️  Rankings scraping error: %v", err)
	}

	// Scrape current tournaments
	if err := s.ScrapeTournaments(ctx); err != nil {
		log.Printf("⚠️  Tournaments scraping error: %v", err)
	}

	// Scrape live scores
	if err := s.ScrapeLiveScores(ctx); err != nil {
		log.Printf("⚠️  Live scores scraping error: %v", err)
	}

	log.Println("✅ ATP Tour scraping complete!")
	return nil
}

// ScrapeRankings scrapes current ATP rankings from atptour.com
func (s *ATPTourScraper) ScrapeRankings(ctx context.Context) error {
	log.Println("📊 Scraping ATP rankings from atptour.com...")

	req, err := http.NewRequestWithContext(ctx, "GET", ATPRankingsPage, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch rankings page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	updateCount := 0

	// Parse rankings table
	doc.Find("table.mega-table tbody tr").Each(func(i int, row *goquery.Selection) {
		// Extract rank, player name, points
		rank := strings.TrimSpace(row.Find("td.rank-cell").Text())
		playerName := strings.TrimSpace(row.Find("td.player-cell a").Text())
		pointsText := strings.TrimSpace(row.Find("td.points-cell").Text())

		if playerName == "" {
			return
		}

		// Clean up data
		rankNum, _ := strconv.Atoi(rank)
		pointsNum, _ := strconv.Atoi(strings.ReplaceAll(pointsText, ",", ""))

		// Generate player ID
		playerID := s.generatePlayerID(playerName)

		// Update or create player
		player, err := s.playerRepo.GetByID(ctx, playerID)
		if err != nil || player == nil {
			// Create new player
			player = &domain.Player{
				ID:          playerID,
				Name:        playerName,
				CountryCode: "XX", // Will be updated by detailed scrape
				Rank:        rankNum,
				Points:      pointsNum,
			}
			if err := s.playerRepo.Create(ctx, player); err != nil {
				if !strings.Contains(err.Error(), "duplicate") {
					log.Printf("Failed to create player %s: %v", playerName, err)
				}
			} else {
				updateCount++
			}
		} else {
			// Update existing player's rank and points
			player.Rank = rankNum
			player.Points = pointsNum
			// Update logic would go here if we had an Update method
			updateCount++
		}
	})

	log.Printf("✅ Scraped %d player rankings", updateCount)
	return nil
}

// ScrapeTournaments scrapes current and upcoming tournaments
func (s *ATPTourScraper) ScrapeTournaments(ctx context.Context) error {
	log.Println("🏆 Scraping ATP tournaments from atptour.com...")

	req, err := http.NewRequestWithContext(ctx, "GET", ATPTournamentsPage, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch tournaments page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	updateCount := 0

	// Parse tournament listings
	doc.Find(".tournament-item, .tourney-result").Each(func(i int, item *goquery.Selection) {
		tourneyName := strings.TrimSpace(item.Find(".tourney-title, h3").Text())
		location := strings.TrimSpace(item.Find(".tourney-location").Text())
		surface := strings.TrimSpace(item.Find(".tourney-surface").Text())
		dates := strings.TrimSpace(item.Find(".tourney-dates").Text())

		if tourneyName == "" {
			return
		}

		// Parse tournament data and create/update
		tourneyID := s.generateTournamentID(tourneyName, time.Now().Year())

		tournament := &domain.Tournament{
			ID:      tourneyID,
			Name:    tourneyName,
			Surface: surface,
			City:    location,
			Year:    time.Now().Year(),
			Status:  "upcoming",
		}

		// Check if exists
		existing, err := s.tournamentRepo.GetByID(ctx, tourneyID)
		if err != nil || existing == nil {
			if err := s.tournamentRepo.Create(ctx, tournament); err != nil {
				if !strings.Contains(err.Error(), "duplicate") {
					log.Printf("Failed to create tournament %s: %v", tourneyName, err)
				}
			} else {
				updateCount++
			}
		}
	})

	log.Printf("✅ Scraped %d tournaments", updateCount)
	return nil
}

// ScrapeLiveScores scrapes live match scores
func (s *ATPTourScraper) ScrapeLiveScores(ctx context.Context) error {
	log.Println("🎾 Scraping live scores from atptour.com...")

	req, err := http.NewRequestWithContext(ctx, "GET", ATPScoresPage, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch scores page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	updateCount := 0

	// Parse live matches
	doc.Find(".match-item, .day-table tbody tr").Each(func(i int, match *goquery.Selection) {
		player1 := strings.TrimSpace(match.Find(".player-left, .player1").Text())
		player2 := strings.TrimSpace(match.Find(".player-right, .player2").Text())
		score := strings.TrimSpace(match.Find(".score").Text())
		status := strings.TrimSpace(match.Find(".status").Text())

		if player1 == "" || player2 == "" {
			return
		}

		// Create or update match
		player1ID := s.generatePlayerID(player1)
		player2ID := s.generatePlayerID(player2)
		matchID := fmt.Sprintf("live-%s-vs-%s-%d", player1ID, player2ID, time.Now().Unix())

		matchStatus := domain.StatusScheduled
		if strings.Contains(strings.ToLower(status), "live") || strings.Contains(strings.ToLower(status), "in progress") {
			matchStatus = domain.StatusLive
		} else if strings.Contains(strings.ToLower(status), "finished") || strings.Contains(strings.ToLower(status), "completed") {
			matchStatus = domain.StatusFinished
		}

		matchObj := &domain.Match{
			ID:          matchID,
			Player1ID:   player1ID,
			Player2ID:   player2ID,
			Status:      matchStatus,
			StartTime:   time.Now(),
			IsSimulated: false,
		}

		// Check if exists
		existing, err := s.matchRepo.GetByID(ctx, matchID)
		if err != nil || existing == nil {
			if err := s.matchRepo.Create(ctx, matchObj); err != nil {
				if !strings.Contains(err.Error(), "duplicate") && !strings.Contains(err.Error(), "foreign key") {
					log.Printf("Failed to create match %s: %v", matchID, err)
				}
			} else {
				updateCount++
			}
		}
	})

	log.Printf("✅ Scraped %d live scores", updateCount)
	return nil
}

// Helper functions
func (s *ATPTourScraper) generatePlayerID(name string) string {
	id := strings.ToLower(name)
	id = strings.ReplaceAll(id, ".", "")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.TrimSpace(id)
	return id
}

func (s *ATPTourScraper) generateTournamentID(name string, year int) string {
	id := strings.ToLower(name)
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.ReplaceAll(id, "'", "")
	id = strings.TrimSpace(id)
	return fmt.Sprintf("%s-%d", id, year)
}
