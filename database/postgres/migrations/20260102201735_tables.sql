-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE game_type AS ENUM ('PICK_FOUR_EVENING', 'PICK_THREE_EVENING');

CREATE TABLE calendar
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       TEXT                     NOT NULL,
    game_type  game_type                NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index for filtering by game_type
CREATE INDEX idx_calendar_game_type ON calendar(game_type);

-- Index for sorting/filtering by created_at
CREATE INDEX idx_calendar_created_at ON calendar(created_at DESC);

-- Composite index for common queries
CREATE INDEX idx_calendar_game_type_created_at ON calendar(game_type, created_at DESC);


CREATE TABLE day_of_calendar
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    calendar_id UUID                     NOT NULL REFERENCES calendar(id) ON DELETE CASCADE,
    date        TIMESTAMP WITH TIME ZONE NOT NULL,
    prize       TEXT                     NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Critical: Foreign key index for JOIN performance
CREATE INDEX idx_day_of_calendar_calendar_id ON day_of_calendar(calendar_id);

-- Index for date queries and sorting
CREATE INDEX idx_day_of_calendar_date ON day_of_calendar(date DESC);

-- Composite index for common queries (calendar + date range)
CREATE INDEX idx_day_of_calendar_calendar_date ON day_of_calendar(calendar_id, date DESC);

-- Index for created_at queries
CREATE INDEX idx_day_of_calendar_created_at ON day_of_calendar(created_at DESC);


CREATE TABLE player
(
    id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

-- Index for searching by player name
CREATE INDEX idx_player_name ON player(name);


CREATE TABLE player_numbers
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_id       UUID                     NOT NULL REFERENCES player(id) ON DELETE CASCADE,
    calendar_id     UUID                     NOT NULL REFERENCES calendar(id) ON DELETE CASCADE,
    winning_numbers TEXT[]                   NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Critical: Foreign key indexes for JOIN performance
CREATE INDEX idx_player_numbers_player_id ON player_numbers(player_id);
CREATE INDEX idx_player_numbers_calendar_id ON player_numbers(calendar_id);

-- Composite index for common queries (player's numbers for a specific calendar)
CREATE INDEX idx_player_numbers_player_calendar ON player_numbers(player_id, calendar_id);

-- Index for created_at queries
CREATE INDEX idx_player_numbers_created_at ON player_numbers(created_at DESC);

-- GIN index for array containment queries on winning_numbers
CREATE INDEX idx_player_numbers_winning_numbers ON player_numbers USING GIN(winning_numbers);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS player_numbers;
DROP TABLE IF EXISTS day_of_calendar;
DROP TABLE IF EXISTS player;
DROP TABLE IF EXISTS calendar;
DROP TYPE IF EXISTS game_type;
-- +goose StatementEnd