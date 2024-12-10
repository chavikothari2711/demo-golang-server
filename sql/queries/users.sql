-- name: CreateUsers :one
INSERT INTO users (id, created_at, updated_at, name,email, api_key)
VALUES ($1, $2, $3, $4, $5, encode(sha256(random()::text::bytea),'hex'))
RETURNING *;


-- name: UpdateUsers :one
UPDATE users 
SET name = $1, email = $2, updated_at = $3 WHERE api_key = $4
RETURNING *;


-- name: GetUsers :one
SELECT * FROM users WHERE email = $1;


-- name: GetUserByAPIKeys :one
SELECT * FROM users WHERE api_key = $1;