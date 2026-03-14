-- +goose Up
CREATE TABLE IF NOT EXISTS links (
	   id INTEGER PRIMARY KEY,
	   long_link TEXT NOT NULL,
	   short_link TEXT NOT NULL
);

-- +goose Down
DROP TABLE links
