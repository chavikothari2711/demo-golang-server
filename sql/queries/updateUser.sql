-- name: UpdateUser :one
UPDATE users 
SET name = $1, email = $2, updated_at = $3 WHERE id = $4
RETURNING *;
