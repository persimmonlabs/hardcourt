CREATE TABLE IF NOT EXISTS tournaments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surface VARCHAR(50) NOT NULL, -- Hard, Clay, Grass
    city VARCHAR(255) NOT NULL,
    country VARCHAR(100),
    start_date DATE,
    end_date DATE,
    year INT, -- Year of tournament for easy filtering
    category VARCHAR(50), -- ATP/WTA, Grand Slam, Masters 1000, etc.
    prize_money BIGINT,
    status VARCHAR(50) DEFAULT 'upcoming', -- upcoming, ongoing, completed
    winner_id VARCHAR(255), -- References players(id) - champion
    runner_up_id VARCHAR(255), -- References players(id) - finalist
    logo_url VARCHAR(500), -- Tournament logo/branding
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for tournaments
CREATE INDEX IF NOT EXISTS idx_tournaments_year ON tournaments(year DESC);
CREATE INDEX IF NOT EXISTS idx_tournaments_status_year ON tournaments(status, year DESC);
CREATE INDEX IF NOT EXISTS idx_tournaments_category ON tournaments(category);

CREATE TABLE IF NOT EXISTS players (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country_code VARCHAR(3) NOT NULL,
    rank INT NOT NULL,
    points INT DEFAULT 0,
    age INT,
    height_cm INT,
    weight_kg INT,
    plays VARCHAR(20), -- Right-handed, Left-handed
    backhand VARCHAR(20), -- One-handed, Two-handed
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS matches (
    id VARCHAR(255) PRIMARY KEY,
    tournament_id VARCHAR(255) REFERENCES tournaments(id),
    player1_id VARCHAR(255) REFERENCES players(id),
    player2_id VARCHAR(255) REFERENCES players(id),
    status VARCHAR(50) NOT NULL, -- Scheduled, Live, Finished
    round VARCHAR(50), -- R128, R64, R32, R16, QF, SF, F
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE,
    winner_id VARCHAR(255) REFERENCES players(id),
    duration_minutes INT,
    court VARCHAR(100),
    is_simulated BOOLEAN DEFAULT FALSE, -- TRUE for simulator matches, FALSE for real matches
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for matches
CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status);
CREATE INDEX IF NOT EXISTS idx_matches_tournament ON matches(tournament_id);
CREATE INDEX IF NOT EXISTS idx_matches_start_time ON matches(start_time DESC);
CREATE INDEX IF NOT EXISTS idx_matches_simulated ON matches(is_simulated) WHERE is_simulated = TRUE;

CREATE TABLE IF NOT EXISTS match_stats (
    match_id VARCHAR(255) PRIMARY KEY REFERENCES matches(id),
    aces_p1 INT DEFAULT 0,
    aces_p2 INT DEFAULT 0,
    df_p1 INT DEFAULT 0,
    df_p2 INT DEFAULT 0,
    break_points_p1 INT DEFAULT 0,
    break_points_p2 INT DEFAULT 0,
    winners_p1 INT DEFAULT 0,
    winners_p2 INT DEFAULT 0,
    unforced_errors_p1 INT DEFAULT 0,
    unforced_errors_p2 INT DEFAULT 0,
    first_serve_pct_p1 FLOAT DEFAULT 0,
    first_serve_pct_p2 FLOAT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Historical score storage for set-by-set data
CREATE TABLE IF NOT EXISTS match_sets (
    id SERIAL PRIMARY KEY,
    match_id VARCHAR(255) REFERENCES matches(id),
    set_number INT NOT NULL,
    games_p1 INT NOT NULL,
    games_p2 INT NOT NULL,
    tiebreak_p1 INT,
    tiebreak_p2 INT,
    UNIQUE(match_id, set_number)
);

-- Tournament draws/brackets
CREATE TABLE IF NOT EXISTS tournament_draws (
    id SERIAL PRIMARY KEY,
    tournament_id VARCHAR(255) REFERENCES tournaments(id),
    round VARCHAR(50) NOT NULL,
    position INT NOT NULL,
    player_id VARCHAR(255) REFERENCES players(id),
    seed INT,
    bye BOOLEAN DEFAULT false,
    UNIQUE(tournament_id, round, position)
);

-- Match highlights/key moments for historical viewing
CREATE TABLE IF NOT EXISTS match_highlights (
    id SERIAL PRIMARY KEY,
    match_id VARCHAR(255) REFERENCES matches(id),
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    event_type VARCHAR(50), -- break_point, ace, rally, etc.
    description TEXT,
    leverage_index FLOAT
);

CREATE INDEX idx_highlights_match ON match_highlights(match_id);
