# Comprehensive Tournament Data Seeding

This system populates the database with **70+ tournaments**, **50+ players**, **20+ Grand Slam finals**, and **tournament draws** from 2020-2024.

## 📊 Data Overview

### Comprehensive Mode (Default)
- **70+ Tournaments**: Grand Slams, Masters 1000, ATP 500, ATP 250 (2020-2024)
- **50+ Players**: Top 50 ATP rankings + tennis legends with full profiles
- **20+ Match Results**: Complete Grand Slam finals with scores (2020-2024)
- **Tournament Draws**: Bracket seeding with player positions

### Standard Mode
- **38 Tournaments**: Grand Slams and Masters 1000 only
- **14 Players**: Core top 12 + legends
- **20+ Match Results**: Grand Slam finals
- **No Draws**: Basic mode without bracket data

## 🚀 Quick Start

### Comprehensive Seeding (Recommended)

```bash
# Seed everything with comprehensive dataset
go run cmd/seed/main.go --all --comprehensive

# Or use the convenience script
./scripts/seed.sh --all
```

### Selective Seeding

```bash
# Seed only players
go run cmd/seed/main.go --players --comprehensive

# Seed only tournaments
go run cmd/seed/main.go --tournaments --comprehensive

# Seed only match results
go run cmd/seed/main.go --matches

# Seed only tournament draws
go run cmd/seed/main.go --draws

# Standard mode (less data)
go run cmd/seed/main.go --all --comprehensive=false
```

## 📋 Detailed Data Breakdown

### Tournaments by Category

#### Grand Slams (20 tournaments)
- **Australian Open** (2020-2024)
- **Roland Garros** (2020-2024)
- **Wimbledon** (2021-2024, 2020 cancelled)
- **US Open** (2020-2024)

#### Masters 1000 (18 tournaments, 2023-2024)
- BNP Paribas Open (Indian Wells)
- Miami Open
- Monte-Carlo Masters
- Madrid Open
- Italian Open (Rome)
- Canadian Open
- Cincinnati Masters
- Shanghai Masters
- Paris Masters

#### ATP 500 (22 tournaments, 2023-2024)
- ABN AMRO Open (Rotterdam)
- Dubai Tennis Championships
- Barcelona Open
- Queen's Club Championships
- Halle Open
- Hamburg European Open
- Citi Open (Washington)
- China Open (Beijing)
- Japan Open (Tokyo)
- Erste Bank Open (Vienna)
- Swiss Indoors (Basel)

#### ATP 250 (13 major events, 2024)
- Adelaide International
- Open Sud de France (Montpellier)
- Open 13 Provence (Marseille)
- Estoril Open
- BMW Open (Munich)
- Geneva Open
- Eastbourne International
- Hall of Fame Open (Newport)
- Winston-Salem Open
- Chengdu Open
- Stockholm Open
- European Open (Antwerp)
- Moselle Open (Metz)

### Players (50+ total)

#### Top 30 ATP Rankings
1. J. Sinner (IT) - 11,180 pts
2. C. Alcaraz (ES) - 8,500 pts
3. N. Djokovic (RS) - 7,900 pts
4. D. Medvedev (RU) - 5,000 pts
5. A. Zverev (DE) - 4,800 pts
6. A. Rublev (RU) - 4,100 pts
7. H. Rune (DK) - 3,800 pts
8. H. Hurkacz (PL) - 3,500 pts
9. T. Fritz (US) - 3,200 pts
10. S. Tsitsipas (GR) - 3,100 pts

... and 40 more players including:

#### Tennis Legends
- Rafael Nadal (ES)
- Dominic Thiem (AT)
- Nick Kyrgios (AU)
- Matteo Berrettini (IT)
- Gael Monfils (FR)

### Match Results (20+ Grand Slam Finals)

Complete finals data including:
- **Scores**: Set-by-set breakdown
- **Duration**: Match length in minutes
- **Winner**: Tournament champion
- **Date**: Match date

Example matches:
- Australian Open 2024: Sinner def. Medvedev 3-2
- Roland Garros 2024: Alcaraz def. Zverev 3-2
- Wimbledon 2024: Alcaraz def. Djokovic 3-0
- US Open 2024: Sinner def. Fritz 3-0

### Tournament Draws

Bracket data for major tournaments including:
- **Initial draw** positions (R128, R64, R32)
- **Seeded players** with seed numbers
- **Progression** through rounds (SF, F)

Example: Australian Open 2024
- 16 seeded players in initial draw
- Semi-final matchups
- Final pairing

## 💻 Usage Examples

### On Railway

```bash
# SSH into Railway service
railway run bash

# Run comprehensive seeding
cd backend && go run cmd/seed/main.go --all --comprehensive

# Or seed specific components
go run cmd/seed/main.go --tournaments --matches --comprehensive
```

### Local Development

```bash
# Set database URL
export DATABASE_URL="postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"

# Run comprehensive seeding
cd backend
go run cmd/seed/main.go --all --comprehensive

# Expected output:
# ✓ Connected to database
# ✓ Migrations completed
# === Seeding Players ===
# ✓ Player seeding successful!
# === Seeding Tournaments ===
# ✓ Tournament seeding successful!
# === Seeding Match Results ===
# ✓ Match seeding successful!
# === Seeding Tournament Draws ===
# ✓ Draw seeding successful!
# ✓ Seeding complete!
```

## 🏗️ Database Schema

The seeder populates these tables:

- `tournaments` - Tournament information with year, winners
- `players` - Player profiles with rankings, age, height, playing hand
- `matches` - Match results with scores and duration
- `tournament_draws` - Bracket positions and seeding
- `match_stats` - Statistical information per match

## ⚙️ Configuration Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--all` | Seed everything | true |
| `--players` | Seed only players | false |
| `--tournaments` | Seed only tournaments | false |
| `--matches` | Seed only match results | false |
| `--draws` | Seed only tournament draws | false |
| `--comprehensive` | Use comprehensive dataset | true |

## 🔄 Idempotency

The seeder uses `ON CONFLICT DO NOTHING` for safe re-running:
- ✅ No duplicates created
- ✅ Existing records preserved
- ✅ New records added
- ✅ Safe to run multiple times

## 🎯 Player ID Generation

Players are identified by URL-safe IDs:
- "J. Sinner" → `j-sinner`
- "C. Alcaraz" → `c-alcaraz`
- "N. Djokovic" → `n-djokovic`
- "F. Auger-Aliassime" → `f-auger-aliassime`

## 📈 Verification Queries

After seeding, verify the data:

```sql
-- Total counts
SELECT 'tournaments', COUNT(*) FROM tournaments
UNION ALL
SELECT 'players', COUNT(*) FROM players
UNION ALL
SELECT 'matches', COUNT(*) FROM matches
UNION ALL
SELECT 'draws', COUNT(*) FROM tournament_draws;

-- Tournaments by year
SELECT year, COUNT(*), category
FROM tournaments
GROUP BY year, category
ORDER BY year DESC, category;

-- Grand Slam winners by year
SELECT t.name, t.year, p.name as winner
FROM tournaments t
LEFT JOIN players p ON t.winner_id = p.id
WHERE t.category = 'Grand Slam'
ORDER BY t.year DESC, t.start_date;

-- Match results
SELECT
    t.name as tournament,
    t.year,
    m.round,
    p1.name as player1,
    p2.name as player2,
    pw.name as winner,
    m.duration_minutes
FROM matches m
JOIN tournaments t ON m.tournament_id = t.id
JOIN players p1 ON m.player1_id = p1.id
JOIN players p2 ON m.player2_id = p2.id
JOIN players pw ON m.winner_id = pw.id
WHERE m.round = 'F'
ORDER BY t.year DESC;

-- Tournament draw for Australian Open 2024
SELECT
    d.round,
    d.position,
    p.name as player,
    d.seed
FROM tournament_draws d
LEFT JOIN players p ON d.player_id = p.id
WHERE d.tournament_id = 'aus-open-2024'
ORDER BY d.round, d.position;
```

## 🔧 Troubleshooting

### "Failed to connect to database"
- Verify `DATABASE_URL` environment variable
- Check PostgreSQL is running
- Validate connection string format

### "Failed to create tournament"
- Ensure migrations have run
- Check for missing player references
- Verify date formats

### "Match seeding failed"
- Players must be seeded before matches
- Verify tournament IDs exist
- Check player name mappings

### "Draw seeding failed"
- Tournaments must exist before draws
- Players must be seeded
- Verify tournament ID format

## 🚀 Performance

Comprehensive seeding typically takes:
- **Players**: ~2 seconds for 50 players
- **Tournaments**: ~5 seconds for 70 tournaments
- **Matches**: ~1 second for 20 finals
- **Draws**: ~1 second for bracket data
- **Total**: ~10 seconds for complete dataset

## 📦 Architecture

```
backend/
├── cmd/seed/main.go              # CLI entry point
├── internal/seeder/
│   ├── service.go                # Core seeding logic
│   ├── tournament_data.go        # Grand Slams + Masters 1000
│   ├── comprehensive_data.go     # ATP 500 + 250 + Players
│   ├── match_data.go             # Match results
│   └── draw_data.go              # Tournament brackets
└── internal/repository/
    ├── tournament_repository.go
    ├── player_repository.go
    ├── match_repository.go
    └── tournament_draw_repository.go
```

## 🔮 Future Enhancements

Potential expansions:
- [ ] More ATP 250 tournaments (currently 13, could add 30+)
- [ ] WTA tournaments (Grand Slams, WTA 1000)
- [ ] Complete match results for all rounds (not just finals)
- [ ] Player head-to-head records
- [ ] Surface-specific statistics
- [ ] Integration with Jeff Sackmann's tennis_atp repository
- [ ] Automatic updates from ATP website
- [ ] Historical data before 2020

## 📚 Data Sources

Current data is manually curated from:
- Official ATP Tour website
- Wikipedia tennis results
- Historical tournament records

For automated data ingestion, consider:
- [Jeff Sackmann's tennis_atp](https://github.com/JeffSackmann/tennis_atp) - Comprehensive ATP match data
- [ATP Tour API](https://www.atptour.com) - Live tournament data
- [Tennis Abstract](http://www.tennisabstract.com) - Statistical analysis

## 📝 Adding Custom Data

To add more tournaments, edit `backend/internal/seeder/comprehensive_data.go`:

```go
{
    ID: "your-tournament-2025",
    Name: "Your Tournament",
    Surface: "Hard",
    City: "City",
    Country: "Country",
    StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
    EndDate: time.Date(2025, 1, 14, 0, 0, 0, 0, time.UTC),
    Year: 2025,
    Category: "ATP 500",
    PrizeMoney: 2000000,
    Status: "upcoming",
    WinnerName: "", // Empty for upcoming
    RunnerUpName: "",
},
```

---

**Note**: This seeder provides a solid foundation. For production use with thousands of matches, consider integrating with external data sources or APIs for automatic updates.
