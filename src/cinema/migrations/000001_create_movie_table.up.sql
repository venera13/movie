USE movie;
CREATE TABLE IF NOT EXISTS movie (id VARCHAR(40), created_at INT, updated_at INT, deleted_at INT, name VARCHAR(50), description(500));