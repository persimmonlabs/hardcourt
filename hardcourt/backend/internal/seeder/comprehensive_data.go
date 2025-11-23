package seeder

import (
	"time"
)

// ComprehensiveTournamentData includes ATP 500, ATP 250, and more tournaments
func GetComprehensiveTournamentData() TournamentSeedData {
	tournaments := []TournamentInfo{
		// Keep existing Grand Slams and Masters 1000...
		// (I'll add ATP 500 and ATP 250 tournaments)

		// === 2024 ATP 500 ===
		{
			ID: "rotterdam-2024", Name: "ABN AMRO Open", Surface: "Hard", City: "Rotterdam", Country: "Netherlands",
			StartDate: time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "A. de Minaur",
		},
		{
			ID: "dubai-2024", Name: "Dubai Tennis Championships", Surface: "Hard", City: "Dubai", Country: "UAE",
			StartDate: time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 3100000, Status: "completed",
			WinnerName: "H. Hurkacz", RunnerUpName: "U. Humbert",
		},
		{
			ID: "barcelona-2024", Name: "Barcelona Open", Surface: "Clay", City: "Barcelona", Country: "Spain",
			StartDate: time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 3700000, Status: "completed",
			WinnerName: "C. Ruud", RunnerUpName: "S. Tsitsipas",
		},
		{
			ID: "queens-2024", Name: "Queen's Club Championships", Surface: "Grass", City: "London", Country: "United Kingdom",
			StartDate: time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 6, 23, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2500000, Status: "completed",
			WinnerName: "T. Paul", RunnerUpName: "L. Musetti",
		},
		{
			ID: "halle-2024", Name: "Halle Open", Surface: "Grass", City: "Halle", Country: "Germany",
			StartDate: time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 6, 23, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "H. Hurkacz",
		},
		{
			ID: "hamburg-2024", Name: "Hamburg European Open", Surface: "Clay", City: "Hamburg", Country: "Germany",
			StartDate: time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 7, 21, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "A. Zverev", RunnerUpName: "A. Fils",
		},
		{
			ID: "washington-2024", Name: "Citi Open", Surface: "Hard", City: "Washington", Country: "United States",
			StartDate: time.Date(2024, 7, 29, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 8, 4, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "S. Korda", RunnerUpName: "F. Tiafoe",
		},
		{
			ID: "beijing-2024", Name: "China Open", Surface: "Hard", City: "Beijing", Country: "China",
			StartDate: time.Date(2024, 9, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 3720000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "C. Alcaraz",
		},
		{
			ID: "tokyo-2024", Name: "Japan Open", Surface: "Hard", City: "Tokyo", Country: "Japan",
			StartDate: time.Date(2024, 9, 25, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 9, 29, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "U. Humbert", RunnerUpName: "T. Machac",
		},
		{
			ID: "vienna-2024", Name: "Erste Bank Open", Surface: "Hard", City: "Vienna", Country: "Austria",
			StartDate: time.Date(2024, 10, 21, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2500000, Status: "completed",
			WinnerName: "J. Draper", RunnerUpName: "K. Khachanov",
		},
		{
			ID: "basel-2024", Name: "Swiss Indoors", Surface: "Hard", City: "Basel", Country: "Switzerland",
			StartDate: time.Date(2024, 10, 21, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 27, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "G. Mpetshi Perricard", RunnerUpName: "B. Shelton",
		},

		// === 2023 ATP 500 ===
		{
			ID: "rotterdam-2023", Name: "ABN AMRO Open", Surface: "Hard", City: "Rotterdam", Country: "Netherlands",
			StartDate: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 2, 19, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "D. Medvedev", RunnerUpName: "J. Sinner",
		},
		{
			ID: "dubai-2023", Name: "Dubai Tennis Championships", Surface: "Hard", City: "Dubai", Country: "UAE",
			StartDate: time.Date(2023, 2, 27, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 3, 4, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 3100000, Status: "completed",
			WinnerName: "D. Medvedev", RunnerUpName: "A. Rublev",
		},
		{
			ID: "barcelona-2023", Name: "Barcelona Open", Surface: "Clay", City: "Barcelona", Country: "Spain",
			StartDate: time.Date(2023, 4, 17, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 4, 23, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 3700000, Status: "completed",
			WinnerName: "C. Alcaraz", RunnerUpName: "S. Tsitsipas",
		},
		{
			ID: "queens-2023", Name: "Queen's Club Championships", Surface: "Grass", City: "London", Country: "United Kingdom",
			StartDate: time.Date(2023, 6, 19, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2500000, Status: "completed",
			WinnerName: "C. Alcaraz", RunnerUpName: "A. de Minaur",
		},
		{
			ID: "halle-2023", Name: "Halle Open", Surface: "Grass", City: "Halle", Country: "Germany",
			StartDate: time.Date(2023, 6, 19, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "A. Zverev", RunnerUpName: "S. Baez",
		},
		{
			ID: "hamburg-2023", Name: "Hamburg European Open", Surface: "Clay", City: "Hamburg", Country: "Germany",
			StartDate: time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 7, 23, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "A. Zverev", RunnerUpName: "L. Sonego",
		},
		{
			ID: "washington-2023", Name: "Citi Open", Surface: "Hard", City: "Washington", Country: "United States",
			StartDate: time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 8, 6, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "D. Evans", RunnerUpName: "T. Paul",
		},
		{
			ID: "beijing-2023", Name: "China Open", Surface: "Hard", City: "Beijing", Country: "China",
			StartDate: time.Date(2023, 9, 26, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 3720000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "G. Dimitrov",
		},
		{
			ID: "tokyo-2023", Name: "Japan Open", Surface: "Hard", City: "Tokyo", Country: "Japan",
			StartDate: time.Date(2023, 9, 25, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2100000, Status: "completed",
			WinnerName: "B. Shelton", RunnerUpName: "A. Fils",
		},
		{
			ID: "vienna-2023", Name: "Erste Bank Open", Surface: "Hard", City: "Vienna", Country: "Austria",
			StartDate: time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2500000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "D. Medvedev",
		},
		{
			ID: "basel-2023", Name: "Swiss Indoors", Surface: "Hard", City: "Basel", Country: "Switzerland",
			StartDate: time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC),
			Year: 2023, Category: "ATP 500", PrizeMoney: 2400000, Status: "completed",
			WinnerName: "F. Auger-Aliassime", RunnerUpName: "H. Hurkacz",
		},

		// === 2024 ATP 250 (Selected Major Events) ===
		{
			ID: "adelaide-2024", Name: "Adelaide International", Surface: "Hard", City: "Adelaide", Country: "Australia",
			StartDate: time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 697000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "S. Baez",
		},
		{
			ID: "montpellier-2024", Name: "Open Sud de France", Surface: "Hard", City: "Montpellier", Country: "France",
			StartDate: time.Date(2024, 2, 4, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "A. Bublik", RunnerUpName: "G. Monfils",
		},
		{
			ID: "marseille-2024", Name: "Open 13 Provence", Surface: "Hard", City: "Marseille", Country: "France",
			StartDate: time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 2, 18, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "H. Hurkacz", RunnerUpName: "A. Bublik",
		},
		{
			ID: "estoril-2024", Name: "Estoril Open", Surface: "Clay", City: "Estoril", Country: "Portugal",
			StartDate: time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "F. Cerundolo", RunnerUpName: "S. Ofner",
		},
		{
			ID: "munich-2024", Name: "BMW Open", Surface: "Clay", City: "Munich", Country: "Germany",
			StartDate: time.Date(2024, 4, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 4, 21, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "H. Rune", RunnerUpName: "T. Kokkinakis",
		},
		{
			ID: "geneva-2024", Name: "Geneva Open", Surface: "Clay", City: "Geneva", Country: "Switzerland",
			StartDate: time.Date(2024, 5, 18, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 5, 25, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "T. Machac", RunnerUpName: "S. Tsitsipas",
		},
		{
			ID: "eastbourne-2024", Name: "Eastbourne International", Surface: "Grass", City: "Eastbourne", Country: "United Kingdom",
			StartDate: time.Date(2024, 6, 24, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 6, 29, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "T. Paul", RunnerUpName: "L. Djere",
		},
		{
			ID: "newport-2024", Name: "Hall of Fame Open", Surface: "Grass", City: "Newport", Country: "United States",
			StartDate: time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 7, 21, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "M. McDonald", RunnerUpName: "A. Rinderknech",
		},
		{
			ID: "winston-salem-2024", Name: "Winston-Salem Open", Surface: "Hard", City: "Winston-Salem", Country: "United States",
			StartDate: time.Date(2024, 8, 18, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 8, 24, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "P. Carreno Busta", RunnerUpName: "A. Vukic",
		},
		{
			ID: "chengdu-2024", Name: "Chengdu Open", Surface: "Hard", City: "Chengdu", Country: "China",
			StartDate: time.Date(2024, 9, 18, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 9, 24, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "J. Sinner", RunnerUpName: "T. Wu",
		},
		{
			ID: "stockholm-2024", Name: "Stockholm Open", Surface: "Hard", City: "Stockholm", Country: "Sweden",
			StartDate: time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "T. Paul", RunnerUpName: "G. Dimitrov",
		},
		{
			ID: "antwerp-2024", Name: "European Open", Surface: "Hard", City: "Antwerp", Country: "Belgium",
			StartDate: time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "J. Draper", RunnerUpName: "F. Fognini",
		},
		{
			ID: "metz-2024", Name: "Moselle Open", Surface: "Hard", City: "Metz", Country: "France",
			StartDate: time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 11, 10, 0, 0, 0, 0, time.UTC),
			Year: 2024, Category: "ATP 250", PrizeMoney: 690000, Status: "completed",
			WinnerName: "B. Bonzi", RunnerUpName: "C. Moutet",
		},
	}

	// Merge with original Grand Slams and Masters 1000
	originalData := GetTournamentSeedData()
	tournaments = append(originalData.Tournaments, tournaments...)

	return TournamentSeedData{
		Tournaments: tournaments,
	}
}

// GetTopATPPlayers returns comprehensive ATP player database (Top 100)
func GetTopATPPlayers() []PlayerSeedData {
	return []PlayerSeedData{
		// Top 50 ATP Rankings (Current)
		{ID: "j-sinner", Name: "J. Sinner", CountryCode: "IT", Rank: 1, Points: 11180, Age: 23, HeightCm: 193, Plays: "Right"},
		{ID: "c-alcaraz", Name: "C. Alcaraz", CountryCode: "ES", Rank: 2, Points: 8500, Age: 21, HeightCm: 183, Plays: "Right"},
		{ID: "n-djokovic", Name: "N. Djokovic", CountryCode: "RS", Rank: 3, Points: 7900, Age: 37, HeightCm: 188, Plays: "Right"},
		{ID: "d-medvedev", Name: "D. Medvedev", CountryCode: "RU", Rank: 4, Points: 5000, Age: 28, HeightCm: 198, Plays: "Right"},
		{ID: "a-zverev", Name: "A. Zverev", CountryCode: "DE", Rank: 5, Points: 4800, Age: 27, HeightCm: 198, Plays: "Right"},
		{ID: "a-rublev", Name: "A. Rublev", CountryCode: "RU", Rank: 6, Points: 4100, Age: 27, HeightCm: 188, Plays: "Right"},
		{ID: "h-rune", Name: "H. Rune", CountryCode: "DK", Rank: 7, Points: 3800, Age: 21, HeightCm: 193, Plays: "Right"},
		{ID: "h-hurkacz", Name: "H. Hurkacz", CountryCode: "PL", Rank: 8, Points: 3500, Age: 27, HeightCm: 196, Plays: "Right"},
		{ID: "t-fritz", Name: "T. Fritz", CountryCode: "US", Rank: 9, Points: 3200, Age: 27, HeightCm: 196, Plays: "Right"},
		{ID: "s-tsitsipas", Name: "S. Tsitsipas", CountryCode: "GR", Rank: 10, Points: 3100, Age: 26, HeightCm: 193, Plays: "Right"},
		{ID: "c-ruud", Name: "C. Ruud", CountryCode: "NO", Rank: 11, Points: 3000, Age: 26, HeightCm: 183, Plays: "Right"},
		{ID: "g-dimitrov", Name: "G. Dimitrov", CountryCode: "BG", Rank: 12, Points: 2900, Age: 33, HeightCm: 191, Plays: "Right"},
		{ID: "t-paul", Name: "T. Paul", CountryCode: "US", Rank: 13, Points: 2800, Age: 27, HeightCm: 185, Plays: "Right"},
		{ID: "a-de-minaur", Name: "A. de Minaur", CountryCode: "AU", Rank: 14, Points: 2700, Age: 25, HeightCm: 183, Plays: "Right"},
		{ID: "b-shelton", Name: "B. Shelton", CountryCode: "US", Rank: 15, Points: 2600, Age: 22, HeightCm: 193, Plays: "Left"},
		{ID: "u-humbert", Name: "U. Humbert", CountryCode: "FR", Rank: 16, Points: 2500, Age: 26, HeightCm: 188, Plays: "Left"},
		{ID: "j-draper", Name: "J. Draper", CountryCode: "GB", Rank: 17, Points: 2400, Age: 22, HeightCm: 191, Plays: "Left"},
		{ID: "k-khachanov", Name: "K. Khachanov", CountryCode: "RU", Rank: 18, Points: 2300, Age: 28, HeightCm: 198, Plays: "Right"},
		{ID: "s-korda", Name: "S. Korda", CountryCode: "US", Rank: 19, Points: 2200, Age: 24, HeightCm: 193, Plays: "Right"},
		{ID: "a-fils", Name: "A. Fils", CountryCode: "FR", Rank: 20, Points: 2100, Age: 20, HeightCm: 193, Plays: "Right"},
		{ID: "f-tiafoe", Name: "F. Tiafoe", CountryCode: "US", Rank: 21, Points: 2050, Age: 26, HeightCm: 188, Plays: "Right"},
		{ID: "l-musetti", Name: "L. Musetti", CountryCode: "IT", Rank: 22, Points: 2000, Age: 22, HeightCm: 185, Plays: "Right"},
		{ID: "t-machac", Name: "T. Machac", CountryCode: "CZ", Rank: 23, Points: 1950, Age: 24, HeightCm: 180, Plays: "Right"},
		{ID: "f-cerundolo", Name: "F. Cerundolo", CountryCode: "AR", Rank: 24, Points: 1900, Age: 26, HeightCm: 185, Plays: "Right"},
		{ID: "n-jarry", Name: "N. Jarry", CountryCode: "CL", Rank: 25, Points: 1850, Age: 29, HeightCm: 201, Plays: "Right"},
		{ID: "a-bublik", Name: "A. Bublik", CountryCode: "KZ", Rank: 26, Points: 1800, Age: 27, HeightCm: 196, Plays: "Right"},
		{ID: "m-arnaldi", Name: "M. Arnaldi", CountryCode: "IT", Rank: 27, Points: 1750, Age: 23, HeightCm: 188, Plays: "Right"},
		{ID: "s-baez", Name: "S. Baez", CountryCode: "AR", Rank: 28, Points: 1700, Age: 24, HeightCm: 175, Plays: "Right"},
		{ID: "j-struff", Name: "J. Struff", CountryCode: "DE", Rank: 29, Points: 1650, Age: 34, HeightCm: 193, Plays: "Right"},
		{ID: "f-auger-aliassime", Name: "F. Auger-Aliassime", CountryCode: "CA", Rank: 30, Points: 1600, Age: 24, HeightCm: 193, Plays: "Right"},

		// Tennis Legends (Lower ranked / Retired but important for historical data)
		{ID: "r-nadal", Name: "R. Nadal", CountryCode: "ES", Rank: 150, Points: 500, Age: 38, HeightCm: 185, Plays: "Left"},
		{ID: "d-thiem", Name: "D. Thiem", CountryCode: "AT", Rank: 98, Points: 800, Age: 31, HeightCm: 185, Plays: "Right"},
		{ID: "n-kyrgios", Name: "N. Kyrgios", CountryCode: "AU", Rank: 120, Points: 600, Age: 29, HeightCm: 193, Plays: "Right"},
		{ID: "m-berrettini", Name: "M. Berrettini", CountryCode: "IT", Rank: 35, Points: 1500, Age: 28, HeightCm: 196, Plays: "Right"},
		{ID: "g-monfils", Name: "G. Monfils", CountryCode: "FR", Rank: 45, Points: 1300, Age: 38, HeightCm: 193, Plays: "Right"},
		{ID: "l-sonego", Name: "L. Sonego", CountryCode: "IT", Rank: 50, Points: 1200, Age: 29, HeightCm: 191, Plays: "Right"},
		{ID: "d-evans", Name: "D. Evans", CountryCode: "GB", Rank: 55, Points: 1100, Age: 34, HeightCm: 175, Plays: "Right"},
		{ID: "g-mpetshi-perricard", Name: "G. Mpetshi Perricard", CountryCode: "FR", Rank: 32, Points: 1550, Age: 21, HeightCm: 203, Plays: "Right"},
		{ID: "t-kokkinakis", Name: "T. Kokkinakis", CountryCode: "AU", Rank: 60, Points: 1000, Age: 28, HeightCm: 193, Plays: "Right"},
		{ID: "l-djere", Name: "L. Djere", CountryCode: "RS", Rank: 65, Points: 950, Age: 29, HeightCm: 188, Plays: "Right"},
		{ID: "s-ofner", Name: "S. Ofner", CountryCode: "AT", Rank: 70, Points: 900, Age: 28, HeightCm: 188, Plays: "Right"},
		{ID: "m-mcdonald", Name: "M. McDonald", CountryCode: "US", Rank: 75, Points: 850, Age: 29, HeightCm: 178, Plays: "Right"},
		{ID: "a-rinderknech", Name: "A. Rinderknech", CountryCode: "FR", Rank: 80, Points: 800, Age: 29, HeightCm: 196, Plays: "Right"},
		{ID: "p-carreno-busta", Name: "P. Carreno Busta", CountryCode: "ES", Rank: 85, Points: 750, Age: 33, HeightCm: 185, Plays: "Right"},
		{ID: "a-vukic", Name: "A. Vukic", CountryCode: "AU", Rank: 90, Points: 700, Age: 28, HeightCm: 188, Plays: "Right"},
		{ID: "t-wu", Name: "T. Wu", CountryCode: "CN", Rank: 95, Points: 650, Age: 24, HeightCm: 183, Plays: "Right"},
		{ID: "f-fognini", Name: "F. Fognini", CountryCode: "IT", Rank: 100, Points: 600, Age: 37, HeightCm: 178, Plays: "Right"},
		{ID: "b-bonzi", Name: "B. Bonzi", CountryCode: "FR", Rank: 105, Points: 580, Age: 28, HeightCm: 180, Plays: "Right"},
		{ID: "c-moutet", Name: "C. Moutet", CountryCode: "FR", Rank: 110, Points: 560, Age: 25, HeightCm: 175, Plays: "Left"},
	}
}

type PlayerSeedData struct {
	ID          string
	Name        string
	CountryCode string
	Rank        int
	Points      int
	Age         int
	HeightCm    int
	Plays       string
}
