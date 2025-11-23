CREATE TABLE IF NOT EXISTS tournaments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surface VARCHAR(50) NOT NULL, -- Hard, Clay, Grass
    city VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS players (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country_code VARCHAR(3) NOT NULL,
    rank INT NOT NULL
);

CREATE TABLE IF NOT EXISTS matches (
    id VARCHAR(255) PRIMARY KEY,
    tournament_id VARCHAR(255) REFERENCES tournaments(id),
    player1_id VARCHAR(255) REFERENCES players(id),
    player2_id VARCHAR(255) REFERENCES players(id),
    status VARCHAR(50) NOT NULL, -- Scheduled, Live, Finished
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    winner_id VARCHAR(255) REFERENCES players(id)
);

CREATE TABLE IF NOT EXISTS match_stats (
    match_id VARCHAR(255) PRIMARY KEY REFERENCES matches(id),
    aces_p1 INT DEFAULT 0,
    aces_p2 INT DEFAULT 0,
    df_p1 INT DEFAULT 0,
    df_p2 INT DEFAULT 0,
    break_points_p1 INT DEFAULT 0,
    break_points_p2 INT DEFAULT 0
);
