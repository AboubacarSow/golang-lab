-- +goose Up
ALTER TABLE posts ALTER COLUMN feed_id SET NOT NULL;

-- +goose Down
ALTER TABLE posts ALTER COLUMN feed_id DROP NOT NULL;