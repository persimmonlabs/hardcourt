# Tournament Data Seeding

This document explains how to populate the database with historical tournament data (2020-2024).

## Overview

The seeding system populates the database with:
- **38 Major Tournaments** from 2020-2024
  - All Grand Slams (Australian Open, Roland Garros, Wimbledon, US Open)
  - Masters 1000 tournaments (Indian Wells, Miami, Monte Carlo, Madrid, Rome, Canada, Cincinnati, Shanghai, Paris)
- **Core ATP Players** (Top 12 + tennis legends)
- **Tournament Winners and Runners-up** for historical context

## Data Included

### Grand Slams (2020-2024)
- Australian Open (20 tournaments)
- Roland Garros
- Wimbledon (except 2020 - cancelled)
- US Open

### Masters 1000 (2023-2024)
- Indian Wells (BNP Paribas Open)
- Miami Open
- Monte-Carlo Masters
- Madrid Open
- Italian Open (Rome)
- Canadian Open
- Cincinnati Masters
- Shanghai Masters
- Paris Masters

## Running the Seeder

### Local Development

```bash
# Set your database connection
export DATABASE_URL="postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"

# Seed everything (players + tournaments)
go run cmd/seed/main.go --all

# Or use the convenience scripts
./scripts/seed.sh --all              # Linux/Mac
.\scripts\seed.ps1 -Mode "all"       # Windows PowerShell

# Seed only players
go run cmd/seed/main.go --players

# Seed only tournaments
go run cmd/seed/main.go --tournaments
```

### Railway Deployment

The seeder can be run as a one-time job on Railway:

```bash
# SSH into Railway service
railway run bash

# Run the seeder
cd backend && go run cmd/seed/main.go --all
```

Or add to your deployment process:

```bash
# In Procfile or start script
./backend/scripts/seed.sh --all && ./backend/server
```

## Database Schema Requirements

The seeder requires these tables to exist:
- `tournaments` - with columns: id, name, surface, city, country, start_date, end_date, year, category, prize_money, status, winner_id, runner_up_id
- `players` - with columns: id, name, country_code, rank, points

These are created automatically by the migration system.

## Idempotency

The seeder uses `ON CONFLICT DO NOTHING` for safe re-running:
- Running multiple times won't create duplicates
- Existing records are preserved
- New records are added

## Player ID Generation

Player IDs are generated from names:
- "J. Sinner" → `j-sinner`
- "C. Alcaraz" → `c-alcaraz`
- "N. Djokovic" → `n-djokovic`

## Extending the Data

To add more tournaments, edit `backend/internal/seeder/tournament_data.go`:

```go
{
    ID: "your-tournament-2025",
    Name: "Your Tournament",
    Surface: "Hard",
    City: "City Name",
    Country: "Country",
    StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
    EndDate: time.Date(2025, 1, 14, 0, 0, 0, 0, time.UTC),
    Year: 2025,
    Category: "Grand Slam", // or "Masters 1000", "ATP 500", etc.
    PrizeMoney: 50000000,
    Status: "upcoming", // or "ongoing", "completed"
    WinnerName: "Player Name", // Leave empty for upcoming
    RunnerUpName: "Player Name",
},
```

## Verification

After seeding, verify the data:

```sql
-- Check tournament count
SELECT COUNT(*) FROM tournaments;
-- Should return 38

-- Check tournaments by year
SELECT year, COUNT(*)
FROM tournaments
GROUP BY year
ORDER BY year DESC;

-- Check Grand Slam winners
SELECT name, year, winner_id
FROM tournaments
WHERE category = 'Grand Slam'
ORDER BY year DESC, start_date;

-- Check players
SELECT COUNT(*) FROM players;
-- Should return 14+
```

## Architecture

```
backend/
├── cmd/seed/main.go              # CLI entry point
├── internal/seeder/
│   ├── tournament_data.go        # Tournament definitions
│   └── service.go                # Seeding logic
└── scripts/
    ├── seed.sh                   # Unix/Linux script
    └── seed.ps1                  # Windows PowerShell script
```

## Troubleshooting

### "Failed to connect to database"
- Ensure DATABASE_URL is set correctly
- Check PostgreSQL is running
- Verify connection string format

### "Player seeding completed with errors"
- Check for duplicate player IDs
- Verify player data format
- Review player name → ID mapping

### "Tournament seeding completed with errors"
- Check date formats are valid
- Verify winner/runner-up names exist
- Review foreign key constraints

## Future Enhancements

- [ ] Import match results for each tournament
- [ ] Add ATP 500 and ATP 250 tournaments
- [ ] Include WTA tournaments
- [ ] Parse Jeff Sackmann's tennis_atp repository directly
- [ ] Add tournament draw/bracket data
- [ ] Include player statistics and head-to-head records
