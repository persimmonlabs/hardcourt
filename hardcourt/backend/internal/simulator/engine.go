package simulator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"hardcourt/backend/internal/domain"
	"hardcourt/backend/internal/logic"
	"hardcourt/backend/internal/repository"

	"github.com/go-redis/redis/v8"
)

type Engine struct {
	rdb             *redis.Client
	math            *logic.MathEngine
	matches         map[string]*domain.Match
	updateChan      chan *domain.Match
	matchRepo       *repository.MatchRepository
	playerRepo      *repository.PlayerRepository
	tournamentRepo  *repository.TournamentRepository
}

func NewEngine(rdb *redis.Client, updateChan chan *domain.Match, matchRepo *repository.MatchRepository, playerRepo *repository.PlayerRepository, tournamentRepo *repository.TournamentRepository) *Engine {
	return &Engine{
		rdb:            rdb,
		math:           logic.NewMathEngine(),
		matches:        make(map[string]*domain.Match),
		updateChan:     updateChan,
		matchRepo:      matchRepo,
		playerRepo:     playerRepo,
		tournamentRepo: tournamentRepo,
	}
}

func (e *Engine) InitializeMatches() {
	ctx := context.Background()

	// Create 5 dummy matches
	tournaments := []domain.Tournament{
		{ID: "t1", Name: "Australian Open", Surface: "Hard", City: "Melbourne"},
		{ID: "t2", Name: "Roland Garros", Surface: "Clay", City: "Paris"},
	}

	players := []domain.Player{
		{ID: "p1", Name: "J. Sinner", CountryCode: "IT", Rank: 1},
		{ID: "p2", Name: "C. Alcaraz", CountryCode: "ES", Rank: 2},
		{ID: "p3", Name: "N. Djokovic", CountryCode: "RS", Rank: 3},
		{ID: "p4", Name: "D. Medvedev", CountryCode: "RU", Rank: 4},
		{ID: "p5", Name: "A. Zverev", CountryCode: "DE", Rank: 5},
		{ID: "p6", Name: "A. Rublev", CountryCode: "RU", Rank: 6},
		{ID: "p7", Name: "H. Rune", CountryCode: "DK", Rank: 7},
		{ID: "p8", Name: "H. Hurkacz", CountryCode: "PL", Rank: 8},
		{ID: "p9", Name: "T. Fritz", CountryCode: "US", Rank: 9},
		{ID: "p10", Name: "S. Tsitsipas", CountryCode: "GR", Rank: 10},
	}

	// Persist tournaments and players to database
	for _, t := range tournaments {
		if err := e.tournamentRepo.Create(ctx, &t); err != nil {
			log.Printf("Warning: Failed to create tournament %s: %v", t.ID, err)
		}
	}

	for _, p := range players {
		if err := e.playerRepo.Create(ctx, &p); err != nil {
			log.Printf("Warning: Failed to create player %s: %v", p.ID, err)
		}
	}

	// Create matches
	for i := 0; i < 5; i++ {
		mID := fmt.Sprintf("match_%d", i)
		match := &domain.Match{
			ID:           mID,
			TournamentID: tournaments[i%2].ID,
			Player1ID:    players[i*2].ID,
			Player2ID:    players[i*2+1].ID,
			Player1:      &players[i*2],
			Player2:      &players[i*2+1],
			Status:       domain.StatusLive,
			StartTime:    time.Now(),
			Score: domain.ScoreState{
				PointsP1: "0",
				PointsP2: "0",
				Serving:  1,
			},
			Stats:     domain.MatchStats{},
			WinProbP1: 0.5,
		}

		// Store in memory
		e.matches[mID] = match

		// Persist to database
		if err := e.matchRepo.Create(ctx, match); err != nil {
			log.Printf("Warning: Failed to create match %s: %v", mID, err)
		}
	}

	log.Printf("Initialized %d matches", len(e.matches))
}

func (e *Engine) Start(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e.simulateRound()
		}
	}
}

func (e *Engine) simulateRound() {
	for _, m := range e.matches {
		if m.Status == domain.StatusFinished {
			continue
		}

		// Simulate a point
		winner := 1
		if rand.Float64() > 0.5 {
			winner = 2
		}

		// Update Score Logic (Simplified for brevity, but functional)
		e.updateScore(m, winner)

		// Update Stats
		if rand.Float64() < 0.1 {
			if winner == 1 {
				m.Stats.AcesP1++
			} else {
				m.Stats.AcesP2++
			}
		}

		// Rally Simulation
		m.Stats.RallyCount = rand.Intn(15) + 1

		// Math Engine Calculations
		m.WinProbP1 = e.math.CalculateWinProbability(
			m.Score.SetsP1, m.Score.SetsP2,
			m.Score.GamesP1, m.Score.GamesP2,
			m.Score.PointsP1, m.Score.PointsP2,
			m.Score.Serving,
		)

		isBreakPoint := (m.Score.Serving == 1 && m.Score.PointsP2 == "40") || (m.Score.Serving == 2 && m.Score.PointsP1 == "40")
		isSetPoint := (m.Score.GamesP1 == 5 && m.Score.GamesP2 < 5) || (m.Score.GamesP2 == 5 && m.Score.GamesP1 < 5) // Simplified
		isMatchPoint := (m.Score.SetsP1 == 2) || (m.Score.SetsP2 == 2)                                               // Simplified

		m.LeverageIndex = e.math.CalculateLeverage(m.WinProbP1, isBreakPoint, isSetPoint, isMatchPoint)
		m.FatigueP1 = e.math.CalculateFatigue(m.FatigueP1, m.Stats.RallyCount)
		m.FatigueP2 = e.math.CalculateFatigue(m.FatigueP2, m.Stats.RallyCount)

		// Publish to Redis
		data, _ := json.Marshal(m)
		e.rdb.Publish(context.Background(), "live_scores", data)

		// Persist to database
		if err := e.matchRepo.Update(context.Background(), m); err != nil {
			log.Printf("Warning: Failed to update match %s in database: %v", m.ID, err)
		}

		// Send to internal channel for WS
		e.updateChan <- m
	}
}

func (e *Engine) updateScore(m *domain.Match, winner int) {
	// Very basic tennis scoring state machine
	points := []string{"0", "15", "30", "40", "AD"}

	var p1Idx, p2Idx int
	for i, p := range points {
		if m.Score.PointsP1 == p {
			p1Idx = i
		}
		if m.Score.PointsP2 == p {
			p2Idx = i
		}
	}

	if winner == 1 {
		if p1Idx == 3 && p2Idx < 3 {
			// Game P1
			e.winGame(m, 1)
		} else if p1Idx == 3 && p2Idx == 3 {
			m.Score.PointsP1 = "AD"
		} else if p1Idx == 4 {
			e.winGame(m, 1)
		} else if p2Idx == 4 {
			m.Score.PointsP2 = "40" // Deuce
		} else {
			m.Score.PointsP1 = points[p1Idx+1]
		}
	} else {
		if p2Idx == 3 && p1Idx < 3 {
			e.winGame(m, 2)
		} else if p2Idx == 3 && p1Idx == 3 {
			m.Score.PointsP2 = "AD"
		} else if p2Idx == 4 {
			e.winGame(m, 2)
		} else if p1Idx == 4 {
			m.Score.PointsP1 = "40" // Deuce
		} else {
			m.Score.PointsP2 = points[p2Idx+1]
		}
	}
}

func (e *Engine) winGame(m *domain.Match, winner int) {
	m.Score.PointsP1 = "0"
	m.Score.PointsP2 = "0"

	// Switch server
	if m.Score.Serving == 1 {
		m.Score.Serving = 2
	} else {
		m.Score.Serving = 1
	}

	if winner == 1 {
		m.Score.GamesP1++
	} else {
		m.Score.GamesP2++
	}

	// Set Logic (Simplified: first to 6)
	if m.Score.GamesP1 >= 6 && m.Score.GamesP1 >= m.Score.GamesP2+2 {
		m.Score.SetsP1++
		m.Score.GamesP1 = 0
		m.Score.GamesP2 = 0
	} else if m.Score.GamesP2 >= 6 && m.Score.GamesP2 >= m.Score.GamesP1+2 {
		m.Score.SetsP2++
		m.Score.GamesP1 = 0
		m.Score.GamesP2 = 0
	}
}
