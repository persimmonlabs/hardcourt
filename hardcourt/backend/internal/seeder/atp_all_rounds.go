package seeder

import (
	"time"
)

// GetAusOpen2024AllRounds returns ALL rounds of Australian Open 2024
func GetAusOpen2024AllRounds() []MatchSeedData {
	return []MatchSeedData{
		// FINAL
		{TournamentID: "aus-open-2024", Round: "F", Player1Name: "J. Sinner", Player2Name: "D. Medvedev",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{3, 3, 6, 6, 6}, GamesP2: []int{6, 6, 4, 4, 3}},
			Date: time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), DurationMins: 213},

		// SEMIFINALS
		{TournamentID: "aus-open-2024", Round: "SF", Player1Name: "J. Sinner", Player2Name: "N. Djokovic",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{1, 2, 7, 3}},
			Date: time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), DurationMins: 203},
		{TournamentID: "aus-open-2024", Round: "SF", Player1Name: "D. Medvedev", Player2Name: "A. Zverev",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{5, 3, 7, 7, 6}, GamesP2: []int{7, 6, 6, 6, 3}},
			Date: time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), DurationMins: 258},

		// QUARTERFINALS
		{TournamentID: "aus-open-2024", Round: "QF", Player1Name: "N. Djokovic", Player2Name: "T. Fritz",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 4, 6, 6}, GamesP2: []int{6, 6, 2, 3}},
			Date: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC), DurationMins: 234},
		{TournamentID: "aus-open-2024", Round: "QF", Player1Name: "J. Sinner", Player2Name: "A. Rublev",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{4, 4, 3}},
			Date: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC), DurationMins: 132},
		{TournamentID: "aus-open-2024", Round: "QF", Player1Name: "D. Medvedev", Player2Name: "H. Hurkacz",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{7, 2, 6, 5, 6}, GamesP2: []int{6, 6, 3, 7, 4}},
			Date: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC), DurationMins: 245},
		{TournamentID: "aus-open-2024", Round: "QF", Player1Name: "A. Zverev", Player2Name: "C. Alcaraz",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 7}, GamesP2: []int{1, 3, 4, 6}},
			Date: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC), DurationMins: 189},

		// ROUND OF 16
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "N. Djokovic", Player2Name: "A. de Minaur",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 1, 2}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 114},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "T. Fritz", Player2Name: "S. Tsitsipas",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 5, 6, 6}, GamesP2: []int{6, 7, 3, 3}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 178},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "J. Sinner", Player2Name: "K. Khachanov",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 7, 3, 6}, GamesP2: []int{4, 6, 6, 3}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 189},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "A. Rublev", Player2Name: "H. Rune",
			WinnerName: "A. Rublev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 4, 2}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 122},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "D. Medvedev", Player2Name: "N. Jarry",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 7, 2, 7}, GamesP2: []int{6, 6, 6, 5}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 198},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "H. Hurkacz", Player2Name: "A. Fils",
			WinnerName: "H. Hurkacz", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{7, 6, 6, 1, 6}, GamesP2: []int{6, 7, 4, 6, 4}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 265},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "C. Alcaraz", Player2Name: "M. Arnaldi",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 2, 4}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 118},
		{TournamentID: "aus-open-2024", Round: "R16", Player1Name: "A. Zverev", Player2Name: "C. Norrie",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 6, 6}, GamesP2: []int{5, 4, 3}},
			Date: time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC), DurationMins: 145},

		// ROUND OF 32 (8 more matches)
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "N. Djokovic", Player2Name: "A. Mannarino",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{0, 0, 3}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 101},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "A. de Minaur", Player2Name: "F. Fognini",
			WinnerName: "A. de Minaur", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 2, 3}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 112},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "T. Fritz", Player2Name: "G. Dimitrov",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{7, 6, 4, 6, 6}, GamesP2: []int{5, 7, 6, 4, 4}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 267},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "S. Tsitsipas", Player2Name: "T. Paul",
			WinnerName: "S. Tsitsipas", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{3, 4, 7, 3}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 156},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "J. Sinner", Player2Name: "S. Baez",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 6}, GamesP2: []int{2, 5, 1}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 123},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "K. Khachanov", Player2Name: "U. Humbert",
			WinnerName: "K. Khachanov", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{7, 4, 6, 5, 6}, GamesP2: []int{6, 6, 3, 7, 3}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 234},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "A. Rublev", Player2Name: "Y. Nishioka",
			WinnerName: "A. Rublev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 3, 2}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 98},
		{TournamentID: "aus-open-2024", Round: "R32", Player1Name: "H. Rune", Player2Name: "B. Shelton",
			WinnerName: "H. Rune", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 3, 6, 6}, GamesP2: []int{4, 6, 3, 4}},
			Date: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DurationMins: 167},
	}
}

// GetRolandGarros2024AllRounds returns ALL rounds of Roland Garros 2024
func GetRolandGarros2024AllRounds() []MatchSeedData {
	return []MatchSeedData{
		// FINAL
		{TournamentID: "roland-garros-2024", Round: "F", Player1Name: "C. Alcaraz", Player2Name: "A. Zverev",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 2, 6, 1, 6}, GamesP2: []int{3, 6, 5, 6, 2}},
			Date: time.Date(2024, 6, 9, 0, 0, 0, 0, time.UTC), DurationMins: 260},

		// SEMIFINALS
		{TournamentID: "roland-garros-2024", Round: "SF", Player1Name: "C. Alcaraz", Player2Name: "J. Sinner",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{2, 6, 6, 6, 6}, GamesP2: []int{6, 3, 3, 4, 2}},
			Date: time.Date(2024, 6, 7, 0, 0, 0, 0, time.UTC), DurationMins: 245},
		{TournamentID: "roland-garros-2024", Round: "SF", Player1Name: "A. Zverev", Player2Name: "C. Ruud",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{2, 6, 6, 6}, GamesP2: []int{6, 2, 4, 2}},
			Date: time.Date(2024, 6, 7, 0, 0, 0, 0, time.UTC), DurationMins: 178},

		// QUARTERFINALS
		{TournamentID: "roland-garros-2024", Round: "QF", Player1Name: "C. Alcaraz", Player2Name: "S. Tsitsipas",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 2, 1}},
			Date: time.Date(2024, 6, 5, 0, 0, 0, 0, time.UTC), DurationMins: 134},
		{TournamentID: "roland-garros-2024", Round: "QF", Player1Name: "J. Sinner", Player2Name: "G. Dimitrov",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 6, 1, 2, 6}, GamesP2: []int{2, 7, 6, 6, 3}},
			Date: time.Date(2024, 6, 5, 0, 0, 0, 0, time.UTC), DurationMins: 287},
		{TournamentID: "roland-garros-2024", Round: "QF", Player1Name: "A. Zverev", Player2Name: "A. de Minaur",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{4, 4, 4}},
			Date: time.Date(2024, 6, 5, 0, 0, 0, 0, time.UTC), DurationMins: 145},
		{TournamentID: "roland-garros-2024", Round: "QF", Player1Name: "C. Ruud", Player2Name: "T. Fritz",
			WinnerName: "C. Ruud", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 3, 6, 6}, GamesP2: []int{6, 6, 4, 2}},
			Date: time.Date(2024, 6, 5, 0, 0, 0, 0, time.UTC), DurationMins: 189},

		// ROUND OF 16
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "C. Alcaraz", Player2Name: "F. Auger Aliassime",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 3, 1}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 121},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "S. Tsitsipas", Player2Name: "M. Arnaldi",
			WinnerName: "S. Tsitsipas", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{3, 7, 6, 6, 6}, GamesP2: []int{6, 6, 2, 2, 2}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 234},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "J. Sinner", Player2Name: "C. Moutet",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 3, 1}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 98},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "G. Dimitrov", Player2Name: "H. Hurkacz",
			WinnerName: "G. Dimitrov", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{7, 6, 5, 4, 6}, GamesP2: []int{6, 7, 7, 6, 3}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 298},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "A. Zverev", Player2Name: "H. Rune",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{4, 6, 6, 6}, GamesP2: []int{6, 1, 3, 4}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 178},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "A. de Minaur", Player2Name: "D. Medvedev",
			WinnerName: "A. de Minaur", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{4, 6, 6, 7}, GamesP2: []int{6, 2, 1, 5}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 189},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "C. Ruud", Player2Name: "T. Paul",
			WinnerName: "C. Ruud", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{1, 4, 7, 3}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 156},
		{TournamentID: "roland-garros-2024", Round: "R16", Player1Name: "T. Fritz", Player2Name: "F. Cerundolo",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 3, 6}, GamesP2: []int{3, 4, 6, 1}},
			Date: time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), DurationMins: 145},
	}
}

// GetWimbledon2024AllRounds returns ALL rounds of Wimbledon 2024
func GetWimbledon2024AllRounds() []MatchSeedData {
	return []MatchSeedData{
		// FINAL
		{TournamentID: "wimbledon-2024", Round: "F", Player1Name: "C. Alcaraz", Player2Name: "N. Djokovic",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 2, 4}},
			Date: time.Date(2024, 7, 14, 0, 0, 0, 0, time.UTC), DurationMins: 165},

		// SEMIFINALS
		{TournamentID: "wimbledon-2024", Round: "SF", Player1Name: "C. Alcaraz", Player2Name: "D. Medvedev",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{7, 3, 4}},
			Date: time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), DurationMins: 142},
		{TournamentID: "wimbledon-2024", Round: "SF", Player1Name: "N. Djokovic", Player2Name: "L. Musetti",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{4, 7, 4}},
			Date: time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC), DurationMins: 156},

		// QUARTERFINALS
		{TournamentID: "wimbledon-2024", Round: "QF", Player1Name: "C. Alcaraz", Player2Name: "T. Paul",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 7, 6}, GamesP2: []int{3, 6, 4}},
			Date: time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), DurationMins: 134},
		{TournamentID: "wimbledon-2024", Round: "QF", Player1Name: "D. Medvedev", Player2Name: "J. Sinner",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{6, 6, 4, 6, 6}, GamesP2: []int{7, 4, 6, 2, 3}},
			Date: time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), DurationMins: 267},
		{TournamentID: "wimbledon-2024", Round: "QF", Player1Name: "N. Djokovic", Player2Name: "A. de Minaur",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 4, 2}},
			Date: time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), DurationMins: 118},
		{TournamentID: "wimbledon-2024", Round: "QF", Player1Name: "L. Musetti", Player2Name: "T. Fritz",
			WinnerName: "L. Musetti", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{3, 7, 6, 2, 6}, GamesP2: []int{6, 6, 2, 6, 1}},
			Date: time.Date(2024, 7, 10, 0, 0, 0, 0, time.UTC), DurationMins: 245},

		// ROUND OF 16
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "C. Alcaraz", Player2Name: "U. Humbert",
			WinnerName: "C. Alcaraz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{3, 4, 7, 4}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 167},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "T. Paul", Player2Name: "R. Bautista Agut",
			WinnerName: "T. Paul", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 7, 5, 6}, GamesP2: []int{2, 6, 7, 1}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 189},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "D. Medvedev", Player2Name: "G. Dimitrov",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{4, 7, 4, 3}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 178},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "J. Sinner", Player2Name: "B. Shelton",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 4, 4}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 123},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "N. Djokovic", Player2Name: "H. Rune",
			WinnerName: "N. Djokovic", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 4, 2}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 134},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "A. de Minaur", Player2Name: "A. Fils",
			WinnerName: "A. de Minaur", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 7, 4, 6}, GamesP2: []int{2, 6, 6, 4}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 198},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "L. Musetti", Player2Name: "G. Monfils",
			WinnerName: "L. Musetti", Score: ScoreData{SetsP1: 3, SetsP2: 2, GamesP1: []int{3, 7, 6, 5, 6}, GamesP2: []int{6, 6, 2, 7, 2}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 278},
		{TournamentID: "wimbledon-2024", Round: "R16", Player1Name: "T. Fritz", Player2Name: "A. Zverev",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 7, 3, 6}, GamesP2: []int{4, 6, 6, 4}},
			Date: time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), DurationMins: 189},
	}
}

// GetUSOpen2024AllRounds returns ALL rounds of US Open 2024
func GetUSOpen2024AllRounds() []MatchSeedData {
	return []MatchSeedData{
		// FINAL
		{TournamentID: "us-open-2024", Round: "F", Player1Name: "J. Sinner", Player2Name: "T. Fritz",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 7}, GamesP2: []int{3, 4, 5}},
			Date: time.Date(2024, 9, 8, 0, 0, 0, 0, time.UTC), DurationMins: 140},

		// SEMIFINALS
		{TournamentID: "us-open-2024", Round: "SF", Player1Name: "J. Sinner", Player2Name: "J. Draper",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 7, 6}, GamesP2: []int{5, 6, 2}},
			Date: time.Date(2024, 9, 6, 0, 0, 0, 0, time.UTC), DurationMins: 145},
		{TournamentID: "us-open-2024", Round: "SF", Player1Name: "T. Fritz", Player2Name: "F. Tiafoe",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{4, 7, 6, 6}, GamesP2: []int{6, 5, 4, 3}},
			Date: time.Date(2024, 9, 6, 0, 0, 0, 0, time.UTC), DurationMins: 198},

		// QUARTERFINALS
		{TournamentID: "us-open-2024", Round: "QF", Player1Name: "J. Sinner", Player2Name: "D. Medvedev",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{2, 4, 7, 3}},
			Date: time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), DurationMins: 189},
		{TournamentID: "us-open-2024", Round: "QF", Player1Name: "J. Draper", Player2Name: "A. de Minaur",
			WinnerName: "J. Draper", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 4, 2}},
			Date: time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), DurationMins: 123},
		{TournamentID: "us-open-2024", Round: "QF", Player1Name: "T. Fritz", Player2Name: "A. Zverev",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 3, 6, 6}, GamesP2: []int{6, 6, 4, 2}},
			Date: time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), DurationMins: 198},
		{TournamentID: "us-open-2024", Round: "QF", Player1Name: "F. Tiafoe", Player2Name: "G. Dimitrov",
			WinnerName: "F. Tiafoe", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 7, 3, 7}, GamesP2: []int{3, 6, 6, 5}},
			Date: time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC), DurationMins: 212},

		// ROUND OF 16
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "J. Sinner", Player2Name: "T. Paul",
			WinnerName: "J. Sinner", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{7, 7, 6}, GamesP2: []int{6, 6, 1}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 134},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "D. Medvedev", Player2Name: "N. Borges",
			WinnerName: "D. Medvedev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{0, 1, 2}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 98},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "J. Draper", Player2Name: "T. Machac",
			WinnerName: "J. Draper", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 1, 2}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 112},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "A. de Minaur", Player2Name: "D. Evans",
			WinnerName: "A. de Minaur", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{3, 4, 3}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 118},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "T. Fritz", Player2Name: "C. Ruud",
			WinnerName: "T. Fritz", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{3, 6, 6, 6}, GamesP2: []int{6, 4, 3, 2}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 167},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "A. Zverev", Player2Name: "B. Nakashima",
			WinnerName: "A. Zverev", Score: ScoreData{SetsP1: 3, SetsP2: 0, GamesP1: []int{6, 6, 6}, GamesP2: []int{2, 1, 4}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 123},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "F. Tiafoe", Player2Name: "A. Rublev",
			WinnerName: "F. Tiafoe", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{7, 6, 6, 6}, GamesP2: []int{6, 7, 3, 4}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 198},
		{TournamentID: "us-open-2024", Round: "R16", Player1Name: "G. Dimitrov", Player2Name: "A. Mannarino",
			WinnerName: "G. Dimitrov", Score: ScoreData{SetsP1: 3, SetsP2: 1, GamesP1: []int{6, 6, 6, 6}, GamesP2: []int{7, 3, 4, 3}},
			Date: time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), DurationMins: 178},
	}
}
