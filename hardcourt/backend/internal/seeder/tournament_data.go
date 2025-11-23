package seeder

import (
	"time"
)

// TournamentSeedData contains historical tournament information for seeding
type TournamentSeedData struct {
	Tournaments []TournamentInfo
}

type TournamentInfo struct {
	ID         string
	Name       string
	Surface    string
	City       string
	Country    string
	StartDate  time.Time
	EndDate    time.Time
	Year       int
	Category   string
	PrizeMoney int64
	Status     string
	WinnerName string // Will be converted to player ID during seeding
	RunnerUpName string // Will be converted to player ID during seeding
}

// GetTournamentSeedData returns tournament data for 2020-2024
func GetTournamentSeedData() TournamentSeedData {
	return TournamentSeedData{
		Tournaments: []TournamentInfo{
			// === 2024 TOURNAMENTS ===

			// Grand Slams 2024
			{
				ID: "aus-open-2024", Name: "Australian Open", Surface: "Hard", City: "Melbourne", Country: "Australia",
				StartDate: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Grand Slam", PrizeMoney: 86500000, Status: "completed",
				WinnerName: "J. Sinner", RunnerUpName: "D. Medvedev",
			},
			{
				ID: "roland-garros-2024", Name: "Roland Garros", Surface: "Clay", City: "Paris", Country: "France",
				StartDate: time.Date(2024, 5, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 6, 9, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Grand Slam", PrizeMoney: 53478000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "A. Zverev",
			},
			{
				ID: "wimbledon-2024", Name: "Wimbledon", Surface: "Grass", City: "London", Country: "United Kingdom",
				StartDate: time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Grand Slam", PrizeMoney: 50000000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "N. Djokovic",
			},
			{
				ID: "us-open-2024", Name: "US Open", Surface: "Hard", City: "New York", Country: "United States",
				StartDate: time.Date(2024, 8, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Grand Slam", PrizeMoney: 75000000, Status: "completed",
				WinnerName: "J. Sinner", RunnerUpName: "T. Fritz",
			},

			// Masters 1000 2024
			{
				ID: "indian-wells-2024", Name: "BNP Paribas Open", Surface: "Hard", City: "Indian Wells", Country: "United States",
				StartDate: time.Date(2024, 3, 6, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 3, 17, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 9300000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "D. Medvedev",
			},
			{
				ID: "miami-2024", Name: "Miami Open", Surface: "Hard", City: "Miami", Country: "United States",
				StartDate: time.Date(2024, 3, 19, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "J. Sinner", RunnerUpName: "G. Dimitrov",
			},
			{
				ID: "monte-carlo-2024", Name: "Monte-Carlo Masters", Surface: "Clay", City: "Monte Carlo", Country: "Monaco",
				StartDate: time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 4, 14, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 6035485, Status: "completed",
				WinnerName: "S. Tsitsipas", RunnerUpName: "C. Ruud",
			},
			{
				ID: "madrid-2024", Name: "Madrid Open", Surface: "Clay", City: "Madrid", Country: "Spain",
				StartDate: time.Date(2024, 4, 24, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 5, 5, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "A. Rublev", RunnerUpName: "F. Auger-Aliassime",
			},
			{
				ID: "rome-2024", Name: "Italian Open", Surface: "Clay", City: "Rome", Country: "Italy",
				StartDate: time.Date(2024, 5, 8, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 5, 19, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 8100000, Status: "completed",
				WinnerName: "A. Zverev", RunnerUpName: "N. Jarry",
			},
			{
				ID: "canada-2024", Name: "Canadian Open", Surface: "Hard", City: "Montreal", Country: "Canada",
				StartDate: time.Date(2024, 8, 6, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 6900000, Status: "completed",
				WinnerName: "A. Rublev", RunnerUpName: "M. Arnaldi",
			},
			{
				ID: "cincinnati-2024", Name: "Cincinnati Masters", Surface: "Hard", City: "Cincinnati", Country: "United States",
				StartDate: time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 8, 19, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 6800000, Status: "completed",
				WinnerName: "J. Sinner", RunnerUpName: "T. Fritz",
			},
			{
				ID: "shanghai-2024", Name: "Shanghai Masters", Surface: "Hard", City: "Shanghai", Country: "China",
				StartDate: time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 13, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "T. Fritz",
			},
			{
				ID: "paris-2024", Name: "Paris Masters", Surface: "Hard", City: "Paris", Country: "France",
				StartDate: time.Date(2024, 10, 28, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC),
				Year: 2024, Category: "Masters 1000", PrizeMoney: 5950000, Status: "completed",
				WinnerName: "A. Zverev", RunnerUpName: "H. Rune",
			},

			// === 2023 TOURNAMENTS ===

			// Grand Slams 2023
			{
				ID: "aus-open-2023", Name: "Australian Open", Surface: "Hard", City: "Melbourne", Country: "Australia",
				StartDate: time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 1, 29, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Grand Slam", PrizeMoney: 76500000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "S. Tsitsipas",
			},
			{
				ID: "roland-garros-2023", Name: "Roland Garros", Surface: "Clay", City: "Paris", Country: "France",
				StartDate: time.Date(2023, 5, 28, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Grand Slam", PrizeMoney: 49600000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "C. Ruud",
			},
			{
				ID: "wimbledon-2023", Name: "Wimbledon", Surface: "Grass", City: "London", Country: "United Kingdom",
				StartDate: time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Grand Slam", PrizeMoney: 44700000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "N. Djokovic",
			},
			{
				ID: "us-open-2023", Name: "US Open", Surface: "Hard", City: "New York", Country: "United States",
				StartDate: time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 9, 10, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Grand Slam", PrizeMoney: 65000000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "D. Medvedev",
			},

			// Masters 1000 2023
			{
				ID: "indian-wells-2023", Name: "BNP Paribas Open", Surface: "Hard", City: "Indian Wells", Country: "United States",
				StartDate: time.Date(2023, 3, 8, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 3, 19, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "D. Medvedev",
			},
			{
				ID: "miami-2023", Name: "Miami Open", Surface: "Hard", City: "Miami", Country: "United States",
				StartDate: time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 4, 2, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "D. Medvedev", RunnerUpName: "J. Sinner",
			},
			{
				ID: "monte-carlo-2023", Name: "Monte-Carlo Masters", Surface: "Clay", City: "Monte Carlo", Country: "Monaco",
				StartDate: time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 4, 16, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 5950000, Status: "completed",
				WinnerName: "A. Rublev", RunnerUpName: "H. Rune",
			},
			{
				ID: "madrid-2023", Name: "Madrid Open", Surface: "Clay", City: "Madrid", Country: "Spain",
				StartDate: time.Date(2023, 4, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 5, 7, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 8400000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "J. Struff",
			},
			{
				ID: "rome-2023", Name: "Italian Open", Surface: "Clay", City: "Rome", Country: "Italy",
				StartDate: time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 5, 21, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 7800000, Status: "completed",
				WinnerName: "D. Medvedev", RunnerUpName: "H. Rune",
			},
			{
				ID: "canada-2023", Name: "Canadian Open", Surface: "Hard", City: "Toronto", Country: "Canada",
				StartDate: time.Date(2023, 8, 7, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 8, 13, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 6800000, Status: "completed",
				WinnerName: "J. Sinner", RunnerUpName: "A. de Minaur",
			},
			{
				ID: "cincinnati-2023", Name: "Cincinnati Masters", Surface: "Hard", City: "Cincinnati", Country: "United States",
				StartDate: time.Date(2023, 8, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 8, 20, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 6800000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "C. Alcaraz",
			},
			{
				ID: "shanghai-2023", Name: "Shanghai Masters", Surface: "Hard", City: "Shanghai", Country: "China",
				StartDate: time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 8800000, Status: "completed",
				WinnerName: "H. Hurkacz", RunnerUpName: "A. Rublev",
			},
			{
				ID: "paris-2023", Name: "Paris Masters", Surface: "Hard", City: "Paris", Country: "France",
				StartDate: time.Date(2023, 10, 30, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 11, 5, 0, 0, 0, 0, time.UTC),
				Year: 2023, Category: "Masters 1000", PrizeMoney: 5950000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "G. Dimitrov",
			},

			// === 2022 TOURNAMENTS ===

			// Grand Slams 2022
			{
				ID: "aus-open-2022", Name: "Australian Open", Surface: "Hard", City: "Melbourne", Country: "Australia",
				StartDate: time.Date(2022, 1, 17, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2022, 1, 30, 0, 0, 0, 0, time.UTC),
				Year: 2022, Category: "Grand Slam", PrizeMoney: 75000000, Status: "completed",
				WinnerName: "R. Nadal", RunnerUpName: "D. Medvedev",
			},
			{
				ID: "roland-garros-2022", Name: "Roland Garros", Surface: "Clay", City: "Paris", Country: "France",
				StartDate: time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2022, 6, 5, 0, 0, 0, 0, time.UTC),
				Year: 2022, Category: "Grand Slam", PrizeMoney: 43600000, Status: "completed",
				WinnerName: "R. Nadal", RunnerUpName: "C. Ruud",
			},
			{
				ID: "wimbledon-2022", Name: "Wimbledon", Surface: "Grass", City: "London", Country: "United Kingdom",
				StartDate: time.Date(2022, 6, 27, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2022, 7, 10, 0, 0, 0, 0, time.UTC),
				Year: 2022, Category: "Grand Slam", PrizeMoney: 40350000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "N. Kyrgios",
			},
			{
				ID: "us-open-2022", Name: "US Open", Surface: "Hard", City: "New York", Country: "United States",
				StartDate: time.Date(2022, 8, 29, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2022, 9, 11, 0, 0, 0, 0, time.UTC),
				Year: 2022, Category: "Grand Slam", PrizeMoney: 60000000, Status: "completed",
				WinnerName: "C. Alcaraz", RunnerUpName: "C. Ruud",
			},

			// === 2021 TOURNAMENTS ===

			// Grand Slams 2021
			{
				ID: "aus-open-2021", Name: "Australian Open", Surface: "Hard", City: "Melbourne", Country: "Australia",
				StartDate: time.Date(2021, 2, 8, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, 2, 21, 0, 0, 0, 0, time.UTC),
				Year: 2021, Category: "Grand Slam", PrizeMoney: 80000000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "D. Medvedev",
			},
			{
				ID: "roland-garros-2021", Name: "Roland Garros", Surface: "Clay", City: "Paris", Country: "France",
				StartDate: time.Date(2021, 5, 30, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, 6, 13, 0, 0, 0, 0, time.UTC),
				Year: 2021, Category: "Grand Slam", PrizeMoney: 38000000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "S. Tsitsipas",
			},
			{
				ID: "wimbledon-2021", Name: "Wimbledon", Surface: "Grass", City: "London", Country: "United Kingdom",
				StartDate: time.Date(2021, 6, 28, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, 7, 11, 0, 0, 0, 0, time.UTC),
				Year: 2021, Category: "Grand Slam", PrizeMoney: 35000000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "M. Berrettini",
			},
			{
				ID: "us-open-2021", Name: "US Open", Surface: "Hard", City: "New York", Country: "United States",
				StartDate: time.Date(2021, 8, 30, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, 9, 12, 0, 0, 0, 0, time.UTC),
				Year: 2021, Category: "Grand Slam", PrizeMoney: 57500000, Status: "completed",
				WinnerName: "D. Medvedev", RunnerUpName: "N. Djokovic",
			},

			// === 2020 TOURNAMENTS ===

			// Grand Slams 2020
			{
				ID: "aus-open-2020", Name: "Australian Open", Surface: "Hard", City: "Melbourne", Country: "Australia",
				StartDate: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC),
				Year: 2020, Category: "Grand Slam", PrizeMoney: 71000000, Status: "completed",
				WinnerName: "N. Djokovic", RunnerUpName: "D. Thiem",
			},
			{
				ID: "roland-garros-2020", Name: "Roland Garros", Surface: "Clay", City: "Paris", Country: "France",
				StartDate: time.Date(2020, 9, 27, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 10, 11, 0, 0, 0, 0, time.UTC),
				Year: 2020, Category: "Grand Slam", PrizeMoney: 38000000, Status: "completed",
				WinnerName: "R. Nadal", RunnerUpName: "N. Djokovic",
			},
			{
				ID: "us-open-2020", Name: "US Open", Surface: "Hard", City: "New York", Country: "United States",
				StartDate: time.Date(2020, 8, 31, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC),
				Year: 2020, Category: "Grand Slam", PrizeMoney: 53400000, Status: "completed",
				WinnerName: "D. Thiem", RunnerUpName: "A. Zverev",
			},
		},
	}
}
