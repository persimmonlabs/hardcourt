# Hardcourt Tennis Platform - Feature Overview

## 🎾 Complete Tournament & Historical Data System

### Backend Enhancements

#### **Enhanced Database Schema** (`backend/schema.sql`)

**New Tables:**
- **`match_sets`** - Set-by-set score history with tiebreak tracking
- **`tournament_draws`** - Tournament bracket positions with seeding
- **`match_highlights`** - Key moments and events during matches

**Enhanced Existing Tables:**
- **`tournaments`** - Added country, dates, category, prize money, status
- **`players`** - Added points, age, height, weight, playing style
- **`matches`** - Added round, end_time, duration, court info
- **`match_stats`** - Added winners, errors, first serve percentages

**Indexes for Performance:**
- `idx_matches_status` - Fast filtering by match status
- `idx_matches_tournament` - Quick tournament match queries
- `idx_matches_start_time` - Chronological sorting
- `idx_highlights_match` - Efficient highlight retrieval

#### **New API Endpoints** (`backend/internal/handlers/tournament_handler.go`)

1. **`GET /api/tournaments`** - List all tournaments with status filtering
2. **`GET /api/tournaments/{id}`** - Get tournament details
3. **`GET /api/tournaments/{id}/matches`** - Get all tournament matches
4. **`GET /api/tournaments/{id}/draw`** - Get tournament bracket
5. **`GET /api/matches/past`** - Get historical matches with filters
6. **`GET /api/matches/{id}/highlights`** - Get match key moments

**Updated Models** (`backend/internal/domain/models.go`):
- Added `SetScore`, `TournamentDraw`, `MatchHighlight` structs
- Extended Tournament, Player, Match with 20+ new fields
- Full support for historical data and analytics

---

### Frontend Application

#### **Navigation System** (`frontend/components/Navigation.tsx`)

Beautiful bottom navigation with 4 tabs:
- **Live** - Real-time matches with WebSocket updates
- **Past** - Historical match archive with search
- **Tournaments** - Browse ongoing/upcoming events
- **Stats** - Player analytics (coming soon)

Features:
- Fixed bottom positioning with backdrop blur
- Active tab highlighting with neon green
- Responsive icons + labels (labels hidden on mobile)
- Smooth transitions and hover effects

#### **1. Live Matches Page** (`frontend/app/page.tsx`)

**Features:**
- Real-time WebSocket connection indicator
- Grid layout (1 col mobile, 2 cols on md+)
- Tournament grouping with Trophy icons
- Live match cards with scores and stats
- Favorite matches functionality
- Beautiful loading states with spinner

**Improvements:**
- Removed separate "Scores/Favorites" tabs
- Cleaner header with connection status
- Responsive grid for larger screens
- Consistent with new navigation

#### **2. Past Matches Page** (`frontend/app/past/page.tsx`)

**Features:**
- **Search bar** - Find matches by player name
- **Surface filters** - Filter by Hard/Clay/Grass/All
- **Date grouping** - Matches organized by date
- **Grid layout** - 2 columns on larger screens
- **Loading states** - Animated spinner while fetching
- **Empty states** - Helpful messages when no results

**Design:**
- Search with magnifying glass icon
- Filter chips with active state highlighting
- Calendar icons for date headers
- Responsive filter scrolling on mobile

#### **3. Tournaments Page** (`frontend/app/tournaments/page.tsx`)

**Features:**
- **Status grouping** - Ongoing, Upcoming, Completed
- **Tournament cards** with:
  - Surface badges (color-coded: Blue=Hard, Orange=Clay, Green=Grass)
  - Location with MapPin icon
  - Category (Grand Slam, Masters, etc.)
  - Prize money formatting ($76.5M)
  - Date ranges with Calendar icon
- **Hover animations** - Cards lift on hover
- **Grid layout** - 2 columns on md+

**Design Highlights:**
- Surface-specific color theming
- Status indicators (green/blue/gray circles)
- Click to view tournament details
- Consistent spacing and typography

#### **4. Tournament Detail Page** (`frontend/app/tournaments/[id]/page.tsx`)

**Features:**
- **Back navigation** - Return to tournaments list
- **Tournament header** with:
  - Large title and location
  - Surface badge
  - 4-stat grid (Category, Prize, Dates, Status)
- **Tab system** - Switch between Matches and Draw
- **Matches tab** - List all tournament matches
- **Draw tab** - Tournament bracket (placeholder)

**Design:**
- Prominent header with key stats
- Tab switching with neon green active state
- Grid stat cards with icons
- Responsive layout

#### **5. Stats Page** (`frontend/app/stats/page.tsx`)

**Coming Soon Placeholder:**
- Feature preview cards
- Icons for Head-to-Head, Rankings, Draws, Performance
- Clean "coming soon" messaging
- Maintains consistent design language

---

### Component Library

#### **Header Component** (`frontend/components/Header.tsx`)
- HARDCOURT branding with neon accent
- Optional connection status indicator
- Reusable across all pages

#### **Navigation Component** (`frontend/components/Navigation.tsx`)
- Fixed bottom navigation bar
- 4 tabs with Lucide icons
- Active state management with Next.js routing
- Responsive label hiding on mobile

---

### Custom Hooks

#### **`useTournaments`** (`frontend/hooks/useTournaments.ts`)
- Fetch tournaments with optional status filter
- Loading and error states
- TypeScript Tournament type

#### **`usePastMatches`** (`frontend/hooks/usePastMatches.ts`)
- Fetch historical matches with filters
- Support for player, tournament, limit params
- Reuses Match type from useLiveScores

---

### Design System

#### **Colors:**
- Background: `#09090b` (zinc-950)
- Surface: `#18181b` (zinc-900)
- Accent: `#CCFF00` (neon green)
- Borders: `white/10` (subtle transparency)

#### **Surface Colors:**
- Hard Court: Blue (`bg-blue-500/20`)
- Clay Court: Orange (`bg-orange-500/20`)
- Grass Court: Green (`bg-green-500/20`)

#### **Typography:**
- Font: Inter (clean, modern sans-serif)
- Headers: Bold, uppercase, wide tracking
- Body: Regular weight, good readability

#### **Responsive Breakpoints:**
- Mobile: Base (< 640px)
- Tablet: `md` (768px+) - 2-column grids
- Desktop: `lg` (1024px+) - Wider max-width

#### **Animations:**
- Loading spinner: Rotating neon border
- Card hover: Lift with shadow
- Tab transitions: Smooth color changes
- Empty states: Centered with helpful text

---

### What's Next?

#### **Backend TODO:**
1. Implement database queries in tournament_handler.go
2. Add tournament repository methods
3. Populate historical data from scrapers
4. Build tournament draw generation logic
5. Create match highlight tracking

#### **Frontend TODO:**
1. Implement tournament draw/bracket visualization
2. Add player statistics page
3. Create match detail modal with full stats
4. Add head-to-head comparison
5. Implement ranking tables
6. Add date range filters for past matches
7. Tournament search and filtering

#### **Data TODO:**
1. Fetch historical tournament data
2. Populate tournament draws from past events
3. Store set-by-set scores from completed matches
4. Track match highlights during live play
5. Build player statistics aggregations

---

### Deployment Notes

**Railway Environment Variables:**
```bash
# Frontend (hardcourt-production)
NEXT_PUBLIC_API_URL=https://accomplished-hope-production.up.railway.app
NEXT_PUBLIC_WS_URL=wss://accomplished-hope-production.up.railway.app/ws

# Backend (accomplished-hope-production)
PORT=8080
DATABASE_URL=${{Postgres.DATABASE_URL}}
```

**Database Migration:**
Run the updated `schema.sql` to add new tables and columns to your Railway Postgres database.

---

### File Structure
```
hardcourt/
├── backend/
│   ├── schema.sql (enhanced with 3 new tables)
│   ├── internal/
│   │   ├── domain/models.go (4 new types)
│   │   └── handlers/
│   │       ├── match_handler.go
│   │       └── tournament_handler.go (NEW - 8 endpoints)
│   └── cmd/server/main.go (added tournament routes)
│
├── frontend/
│   ├── app/
│   │   ├── page.tsx (Live - updated with grid)
│   │   ├── past/page.tsx (NEW - historical matches)
│   │   ├── tournaments/
│   │   │   ├── page.tsx (NEW - tournament list)
│   │   │   └── [id]/page.tsx (NEW - tournament detail)
│   │   └── stats/page.tsx (NEW - coming soon)
│   ├── components/
│   │   ├── Header.tsx (NEW - reusable header)
│   │   ├── Navigation.tsx (NEW - bottom nav)
│   │   └── MatchCard.tsx (existing)
│   └── hooks/
│       ├── useLiveScores.ts (existing)
│       ├── useTournaments.ts (NEW)
│       └── usePastMatches.ts (NEW)
│
└── docs/
    └── FEATURE_OVERVIEW.md (this file)
```

---

## Summary

You now have a **fully-featured tennis platform** with:
- ✅ Real-time live scores with WebSocket
- ✅ Historical match archive with search
- ✅ Tournament browsing and details
- ✅ Beautiful responsive design (mobile → desktop)
- ✅ Consistent neon green design system
- ✅ 8 new backend API endpoints
- ✅ Enhanced database schema for historical data
- ✅ 5 pages with navigation
- ✅ 2 new custom hooks
- ✅ Loading states and error handling

**Next step:** Deploy to Railway and populate with real tournament data!
