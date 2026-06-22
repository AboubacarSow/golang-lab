-- name: GetUser :one
SELECT * FROM users
WHERE id=$1 LIMIT 1;

-- name: GetUserByKey :one
SELECT * FROM users 
WHERE api_key=$1 LIMIT 1;

-- name: GetAll :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users(
    id,created_at,updated_at,name, api_key
    )
VALUES (
    $1,$2,$3,$4, encode(sha256(random()::text::bytea),'hex')
    )
RETURNING *;

-- name: UpdateUser :exec
UPDATE users 
    set updated_at=$2,
     name=$3
WHERE id=$1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id=$1;