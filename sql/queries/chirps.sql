-- name: CreateChirp :one
INSERT INTO chirps (user_id, body)
VALUES ($1, $2)
RETURNING *;

-- name: GetChirps :many
SELECT * FROM chirps ORDER BY created_at ASC;

-- name: GetChirp :one
SELECT * FROM chirps WHERE id = $1;