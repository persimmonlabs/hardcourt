package scraper

import (
	"context"
	"log"
	"sync"
	"time"

	"hardcourt/backend/internal/repository"
)

// Scheduler manages periodic web scraping
type Scheduler struct {
	atpScraper        *ATPTourScraper
	flashScoreScraper *FlashScoreScraper
	interval          time.Duration
	ctx               context.Context
	cancel            context.CancelFunc
	wg                sync.WaitGroup
	running           bool
	mu                sync.Mutex
}

// NewScheduler creates a new scraping scheduler
func NewScheduler(
	tournamentRepo *repository.TournamentRepository,
	playerRepo *repository.PlayerRepository,
	matchRepo *repository.MatchRepository,
	interval time.Duration,
) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())

	return &Scheduler{
		atpScraper:        NewATPTourScraper(tournamentRepo, playerRepo, matchRepo),
		flashScoreScraper: NewFlashScoreScraper(matchRepo, playerRepo),
		interval:          interval,
		ctx:               ctx,
		cancel:            cancel,
		running:           false,
	}
}

// Start begins the scraping scheduler
func (s *Scheduler) Start() error {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return nil
	}
	s.running = true
	s.mu.Unlock()

	log.Printf("🕷️  Starting scraper scheduler (interval: %s)", s.interval)

	// Run initial scrape immediately
	s.runScrapeJobs()

	// Start periodic scraping
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		ticker := time.NewTicker(s.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.runScrapeJobs()
			case <-s.ctx.Done():
				log.Println("🛑 Scraper scheduler stopped")
				return
			}
		}
	}()

	return nil
}

// Stop gracefully stops the scheduler
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	log.Println("🛑 Stopping scraper scheduler...")
	s.cancel()
	s.wg.Wait()
	s.running = false
	log.Println("✅ Scraper scheduler stopped")
}

// runScrapeJobs runs all scraping jobs concurrently
func (s *Scheduler) runScrapeJobs() {
	log.Println("🔄 Running scheduled scrape jobs...")

	// Use a timeout context for each scrape cycle
	ctx, cancel := context.WithTimeout(s.ctx, 55*time.Second) // Leave 5s buffer
	defer cancel()

	var wg sync.WaitGroup

	// ATP Tour scraping
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.atpScraper.ScrapeAll(ctx); err != nil {
			log.Printf("⚠️  ATP Tour scraping error: %v", err)
		}
	}()

	// FlashScore scraping
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.flashScoreScraper.ScrapeLiveMatches(ctx); err != nil {
			log.Printf("⚠️  FlashScore scraping error: %v", err)
		}
	}()

	// Wait for all scrapers to complete
	wg.Wait()

	log.Println("✅ Scrape cycle complete")
}

// GetStatus returns the current scheduler status
func (s *Scheduler) GetStatus() map[string]interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	return map[string]interface{}{
		"running":  s.running,
		"interval": s.interval.String(),
	}
}
