package logic

import (
	"math"
)

// MathEngine handles the "Moneyball" statistics
type MathEngine struct{}

func NewMathEngine() *MathEngine {
	return &MathEngine{}
}

// CalculateWinProbability uses a simplified Markov Chain approach
// In a real scenario, this would be a recursive function solving P(Game) -> P(Set) -> P(Match)
// Here we use a deterministic approximation based on score delta and base strength.
func (m *MathEngine) CalculateWinProbability(setsP1, setsP2, gamesP1, gamesP2 int, pointsP1, pointsP2 string, server int) float64 {
	// Base probability (assume 50/50 start)
	prob := 0.5

	// Weight factors
	const (
		SetWeight   = 0.15
		GameWeight  = 0.05
		PointWeight = 0.02
	)

	// Adjust for Sets
	prob += float64(setsP1-setsP2) * SetWeight

	// Adjust for Games
	prob += float64(gamesP1-gamesP2) * GameWeight

	// Adjust for Points (Simplified mapping)
	p1Val := pointToValue(pointsP1)
	p2Val := pointToValue(pointsP2)
	prob += float64(p1Val-p2Val) * PointWeight

	// Server advantage (Server usually has 65% win prob on hard court)
	if server == 1 {
		prob += 0.05
	} else {
		prob -= 0.05
	}

	// Clamp
	if prob > 0.99 {
		return 0.99
	}
	if prob < 0.01 {
		return 0.01
	}
	return prob
}

func pointToValue(p string) int {
	switch p {
	case "0":
		return 0
	case "15":
		return 1
	case "30":
		return 2
	case "40":
		return 3
	case "AD":
		return 4
	default:
		return 0
	}
}

// CalculateLeverage determines the importance of the current point.
// Leverage = |P(WinMatch | WinPoint) - P(WinMatch | LosePoint)|
// High leverage means the outcome of this point swings the match probability significantly.
func (m *MathEngine) CalculateLeverage(currentProb float64, isBreakPoint bool, isSetPoint bool, isMatchPoint bool) float64 {
	baseLeverage := 0.1 // Standard point

	if isBreakPoint {
		baseLeverage += 0.2
	}
	if isSetPoint {
		baseLeverage += 0.3
	}
	if isMatchPoint {
		baseLeverage += 0.4
	}

	// Also, leverage is higher when the match is closer (prob near 0.5)
	uncertainty := 1.0 - 2.0*math.Abs(currentProb-0.5) // 1.0 at 50/50, 0.0 at 100/0

	return baseLeverage * (0.5 + 0.5*uncertainty)
}

// CalculateFatigue Linear decay based on rally count and time
func (m *MathEngine) CalculateFatigue(currentFatigue float64, rallyLength int) float64 {
	// Recovery
	newFatigue := currentFatigue - 0.5
	if newFatigue < 0 {
		newFatigue = 0
	}

	// Exertion
	exertion := float64(rallyLength) * 0.8
	newFatigue += exertion

	if newFatigue > 100 {
		return 100
	}
	return newFatigue
}
