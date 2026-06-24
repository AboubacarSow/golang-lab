-- +goose Up
ALTER TABLE feed_follows ADD CONSTRAINT fk_unique UNIQUE(user_id, feed_id);

-- +goose Down
ALTER TABLE feed_follows DROP CONSTRAINT fk_unique;