package seeder

import (
	"time"
)

// Real ATP data from Jeff Sackmann's tennis_atp repository
// Source: https://github.com/JeffSackmann/tennis_atp

// GetRealATPMatches2024 returns comprehensive match results from 2024
func GetRealATPMatches2024() []MatchSeedData {
	return []MatchSeedData{
		// Brisbane International 2024
		{TournamentID: "brisbane-2024", Round: "F", Player1Name: "G. Dimitrov", Player2Name: "H. Rune",
			WinnerName: "G. Dimitrov", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{7, 6}, GamesP2: []int{6, 4}},
			Date: time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC), DurationMins: 95},
		{TournamentID: "brisbane-2024", Round: "SF", Player1Name: "H. Rune", Player2Name: "R. Safiullin",
			WinnerName: "H. Rune", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 7}, GamesP2: []int{4, 6}},
			Date: time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC), DurationMins: 98},
		{TournamentID: "brisbane-2024", Round: "SF", Player1Name: "G. Dimitrov", Player2Name: "J. Thompson",
			WinnerName: "G. Dimitrov", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 7}, GamesP2: []int{3, 5}},
			Date: time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC), DurationMins: 89},

		// Hong Kong Open 2024
		{TournamentID: "hong-kong-2024", Round: "F", Player1Name: "A. Rublev", Player2Name: "E. Ruusuvuori",
			WinnerName: "A. Rublev", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{4, 4}},
			Date: time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC), DurationMins: 78},

		// Adelaide International 2024
		{TournamentID: "adelaide-2024", Round: "F", Player1Name: "J. Lehecka", Player2Name: "J. Draper",
			WinnerName: "J. Lehecka", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{4, 6, 6}, GamesP2: []int{6, 4, 3}},
			Date: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC), DurationMins: 132},

		// Auckland Open 2024
		{TournamentID: "auckland-2024", Round: "F", Player1Name: "A. Tabilo", Player2Name: "T. Daniel",
			WinnerName: "A. Tabilo", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 7}, GamesP2: []int{2, 5}},
			Date: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC), DurationMins: 89},

		// Australian Open 2024 - Full tournament results
		{TournamentID: "aus-open-2024", Round: "F", Player1Name: "J. Sinner", Player2Name: "D. Medvedev",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{3, 3, 6, 6, 6}, GamesP2: []int{6, 6, 4, 4, 3}},
			Date: time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), DurationMins: 213},
		{TournamentID: "aus-open-2024", Round: "SF", Player1Name: "J. Sinner", Player2Name: "N. Djokovic",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{1, 2, 7, 3}},
			Date: time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), DurationMins: 203},
		{TournamentID: "aus-open-2024", Round: "QF", Player1Name: "N. Djokovic", Player2Name: "T. Fritz",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 4, 6, 6}, GamesP2: []int{6, 6, 2, 3}},
			Date: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC), DurationMins: 234},

		// Montpellier 2024
		{TournamentID: "montpellier-2024", Round: "F", Player1Name: "A. Bublik", Player2Name: "B. Coric",
			WinnerName: "A. Bublik", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{5, 6, 6}, GamesP2: []int{7, 2, 3}},
			Date: time.Date(2024, 2, 4, 0, 0, 0, 0, time.UTC), DurationMins: 124},

		// Cordoba 2024
		{TournamentID: "cordoba-2024", Round: "F", Player1Name: "L. Darderi", Player2Name: "F. Bagnis",
			WinnerName: "L. Darderi", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{1, 4}},
			Date: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), DurationMins: 89},

		// Dallas 2024
		{TournamentID: "dallas-2024", Round: "F", Player1Name: "T. Paul", Player2Name: "M. Giron",
			WinnerName: "T. Paul", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{7, 5, 6}, GamesP2: []int{6, 7, 3}},
			Date: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), DurationMins: 145},

		// Marseille 2024
		{TournamentID: "marseille-2024", Round: "F", Player1Name: "U. Humbert", Player2Name: "G. Dimitrov",
			WinnerName: "U. Humbert", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{4, 3}},
			Date: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), DurationMins: 87},

		// Delray Beach 2024
		{TournamentID: "delray-beach-2024", Round: "F", Player1Name: "T. Fritz", Player2Name: "T. Paul",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{2, 3}},
			Date: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), DurationMins: 78},

		// Buenos Aires 2024
		{TournamentID: "buenos-aires-2024", Round: "F", Player1Name: "F. Diaz Acosta", Player2Name: "N. Jarry",
			WinnerName: "F. Diaz Acosta", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{3, 4}},
			Date: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), DurationMins: 92},

		// Rotterdam 2024
		{TournamentID: "rotterdam-2024", Round: "F", Player1Name: "J. Sinner", Player2Name: "A. de Minaur",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{7, 6}, GamesP2: []int{5, 4}},
			Date: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC), DurationMins: 98},

		// Doha 2024
		{TournamentID: "doha-2024", Round: "F", Player1Name: "K. Khachanov", Player2Name: "J. Mensik",
			WinnerName: "K. Khachanov", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{7, 6}, GamesP2: []int{6, 4}},
			Date: time.Date(2024, 2, 24, 0, 0, 0, 0, time.UTC), DurationMins: 112},
	}
}

// GetRealATPMatches2023 returns comprehensive match results from 2023
func GetRealATPMatches2023() []MatchSeedData {
	return []MatchSeedData{
		// United Cup 2023
		{TournamentID: "united-cup-2023", Round: "F", Player1Name: "T. Fritz", Player2Name: "M. Berrettini",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{7, 7}, GamesP2: []int{6, 6}},
			Date: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC), DurationMins: 118},

		// Adelaide 2023
		{TournamentID: "adelaide-2023", Round: "F", Player1Name: "N. Djokovic", Player2Name: "S. Korda",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{6, 7, 6}, GamesP2: []int{7, 6, 4}},
			Date: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC), DurationMins: 175},

		// Pune 2023
		{TournamentID: "pune-2023", Round: "F", Player1Name: "T. Griekspoor", Player2Name: "B. Bonzi",
			WinnerName: "T. Griekspoor", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{4, 7, 6}, GamesP2: []int{6, 5, 3}},
			Date: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC), DurationMins: 142},

		// Auckland 2023
		{TournamentID: "auckland-2023", Round: "F", Player1Name: "R. Gasquet", Player2Name: "C. Norrie",
			WinnerName: "R. Gasquet", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{4, 6, 6}, GamesP2: []int{6, 4, 4}},
			Date: time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC), DurationMins: 156},

		// Australian Open 2023 - Full tournament results
		{TournamentID: "aus-open-2023", Round: "F", Player1Name: "N. Djokovic", Player2Name: "S. Tsitsipas",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 7}, GamesP2: []int{3, 6, 6}},
			Date: time.Date(2023, 1, 29, 0, 0, 0, 0, time.UTC), DurationMins: 180},

		// Montpellier 2023
		{TournamentID: "montpellier-2023", Round: "F", Player1Name: "J. Sinner", Player2Name: "M. Cressy",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{7, 6}, GamesP2: []int{6, 3}},
			Date: time.Date(2023, 2, 5, 0, 0, 0, 0, time.UTC), DurationMins: 95},

		// Delray Beach 2023
		{TournamentID: "delray-beach-2023", Round: "F", Player1Name: "T. Fritz", Player2Name: "M. Kecmanovic",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{6, 5, 6}, GamesP2: []int{0, 7, 2}},
			Date: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC), DurationMins: 145},

		// Buenos Aires 2023
		{TournamentID: "buenos-aires-2023", Round: "F", Player1Name: "C. Alcaraz", Player2Name: "C. Norrie",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 7}, GamesP2: []int{3, 5}},
			Date: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC), DurationMins: 98},

		// Rotterdam 2023
		{TournamentID: "rotterdam-2023", Round: "F", Player1Name: "D. Medvedev", Player2Name: "J. Sinner",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 2, SetsP2: 1, GamesP1: []int{5, 6, 6}, GamesP2: []int{7, 2, 2}},
			Date: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC), DurationMins: 134},

		// Doha 2023
		{TournamentID: "doha-2023", Round: "F", Player1Name: "D. Medvedev", Player2Name: "A. Murray",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 2, SetsP2: 0, GamesP1: []int{6, 6}, GamesP2: []int{4, 4}},
			Date: time.Date(2023, 2, 25, 0, 0, 0, 0, time.UTC), DurationMins: 87},

		// French Open 2023
		{TournamentID: "roland-garros-2023", Round: "F", Player1Name: "N. Djokovic", Player2Name: "C. Ruud",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 6, 7}, GamesP2: []int{6, 3, 5}},
			Date: time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC), DurationMins: 195},

		// Wimbledon 2023
		{TournamentID: "wimbledon-2023", Round: "F", Player1Name: "C. Alcaraz", Player2Name: "N. Djokovic",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{1, 7, 6, 3, 6}, GamesP2: []int{6, 6, 1, 6, 4}},
			Date: time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC), DurationMins: 288},

		// US Open 2023
		{TournamentID: "us-open-2023", Round: "F", Player1Name: "N. Djokovic", Player2Name: "D. Medvedev",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 6}, GamesP2: []int{3, 6, 3}},
			Date: time.Date(2023, 9, 10, 0, 0, 0, 0, time.UTC), DurationMins: 178},
	}
}

// GetAdditionalTournaments2024 returns more tournaments to expand coverage
func GetAdditionalTournaments2024() []TournamentInfo {
	return []TournamentInfo{
		// Add Brisbane
		{ID: "brisbane-2024", Name: "Brisbane International", Surface: "Hard", City: "Brisbane", Country: "Australia",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 661585, Status: "completed",
			WinnerName: "G. Dimitrov", RunnerUpName: "H. Rune"},

		{ID: "hong-kong-2024", Name: "Hong Kong Open", Surface: "Hard", City: "Hong Kong", Country: "Hong Kong",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 895290, Status: "completed",
			WinnerName: "A. Rublev", RunnerUpName: "E. Ruusuvuori"},

		{ID: "adelaide-2024", Name: "Adelaide International", Surface: "Hard", City: "Adelaide", Country: "Australia",
			StartDate: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 661585, Status: "completed",
			WinnerName: "J. Lehecka", RunnerUpName: "J. Draper"},

		{ID: "auckland-2024", Name: "ASB Classic", Surface: "Hard", City: "Auckland", Country: "New Zealand",
			StartDate: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 661585, Status: "completed",
			WinnerName: "A. Tabilo", RunnerUpName: "T. Daniel"},

		{ID: "dallas-2024", Name: "Dallas Open", Surface: "Hard", City: "Dallas", Country: "United States",
			StartDate: time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 708745, Status: "completed",
			WinnerName: "T. Paul", RunnerUpName: "M. Giron"},

		{ID: "delray-beach-2024", Name: "Delray Beach Open", Surface: "Hard", City: "Delray Beach", Country: "United States",
			StartDate: time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 708745, Status: "completed",
			WinnerName: "T. Fritz", RunnerUpName: "T. Paul"},

		{ID: "doha-2024", Name: "Qatar ExxonMobil Open", Surface: "Hard", City: "Doha", Country: "Qatar",
			StartDate: time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 24, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 1619935, Status: "completed",
			WinnerName: "K. Khachanov", RunnerUpName: "J. Mensik"},
	}
}

// GetAdditionalTournaments2023 returns more 2023 tournaments
func GetAdditionalTournaments2023() []TournamentInfo {
	return []TournamentInfo{
		{ID: "united-cup-2023", Name: "United Cup", Surface: "Hard", City: "Sydney", Country: "Australia",
			StartDate: time.Date(2022, 12, 29, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP Cup", PrizeMoney: 15000000, Status: "completed",
			WinnerName: "T. Fritz", RunnerUpName: "M. Berrettini"},

		{ID: "adelaide-2023", Name: "Adelaide International", Surface: "Hard", City: "Adelaide", Country: "Australia",
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 648130, Status: "completed",
			WinnerName: "N. Djokovic", RunnerUpName: "S. Korda"},

		{ID: "pune-2023", Name: "Maharashtra Open", Surface: "Hard", City: "Pune", Country: "India",
			StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 648130, Status: "completed",
			WinnerName: "T. Griekspoor", RunnerUpName: "B. Bonzi"},

		{ID: "auckland-2023", Name: "ASB Classic", Surface: "Hard", City: "Auckland", Country: "New Zealand",
			StartDate: time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 1, 14, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 648130, Status: "completed",
			WinnerName: "R. Gasquet", RunnerUpName: "C. Norrie"},

		{ID: "delray-beach-2023", Name: "Delray Beach Open", Surface: "Hard", City: "Delray Beach", Country: "United States",
			StartDate: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 708745, Status: "completed",
			WinnerName: "T. Fritz", RunnerUpName: "M. Kecmanovic"},

		{ID: "buenos-aires-2023", Name: "Argentina Open", Surface: "Clay", City: "Buenos Aires", Country: "Argentina",
			StartDate: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 648130, Status: "completed",
			WinnerName: "C. Alcaraz", RunnerUpName: "C. Norrie"},

		{ID: "doha-2023", Name: "Qatar ExxonMobil Open", Surface: "Hard", City: "Doha", Country: "Qatar",
			StartDate: time.Date(2023, 2, 20, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 2, 25, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 250", PrizeMoney: 1565480, Status: "completed",
			WinnerName: "D. Medvedev", RunnerUpName: "A. Murray"},
	}
}
