DROP TABLE IF EXISTS position;
DROP TABLE IF EXISTS player;
DROP TABLE IF EXISTS team_config;

CREATE TABLE position (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE player (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    team TEXT NOT NULL,
    position_id INTEGER NOT NULL,
    FOREIGN KEY (position_id) REFERENCES position (id)
);

CREATE TABLE team_config (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    num_guards INT NOT NULL DEFAULT 0,
    num_wings INT NOT NULL DEFAULT 0,
    num_bigs INT NOT NULL DEFAULT 0,
    player_count INT GENERATED ALWAYS AS (num_guards + num_wings + num_bigs) STORED
);

-- Seed in positions
INSERT INTO position (name)
VALUES
    ('guard'),
    ('wing'),
    ('big');

-- Seed in configs
INSERT INTO team_config (name, num_guards, num_wings, num_bigs)
VALUES
    ('Guards', 1, 0, 0),
    ('Wings', 0, 1, 0),
    ('Bigs', 0, 0, 1),
    ('Two Guards', 2, 0, 0),
    ('Guard and Wing', 1, 1, 0),
    ('Guard and Big', 1, 0, 1),
    ('Two Wings', 0, 2, 0),
    ('Wing and Big', 0, 1, 1),
    ('One Guard, One Wing and One Big', 1, 1, 1),
    ('Three Guards', 3, 0, 0),
    ('Standard 5v5', 1, 2, 2);