-- name: CreateFeed :one
INSERT INTO feeds(
    id,created_at, updated_at,name, url, user_id
)VALUES(
    $1,$2,$3,$4,$5,$6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;


-- name: GetOneFeed :one
SELECT * FROM feeds 
WHERE id=$1 LIMIT 1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds 
ORDER BY lastfetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkAsFetched :one
UPDATE  feeds 
SET updated_at=Now(), 
    lastfetched_at= Now()
WHERE id=$1
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id=$1;