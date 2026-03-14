-- +goose Up
INSERT INTO links (long_link, short_link)
VALUES
("https://google.com", "google"),
("https://pornhub.org", "pornhub");

-- +goose Down
SELECT 'down SQL query';
DELETE FROM links WHERE id IN (1,2)
