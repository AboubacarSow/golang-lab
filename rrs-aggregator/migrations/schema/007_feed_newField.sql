-- +goose Up
ALTER TABLE feeds ADD COLUMN lastfetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds DROP COLUMN lastfetched_at;