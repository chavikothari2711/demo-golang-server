-- name: CreateBlogVisibilityType :one
INSERT INTO blogType (visibilityType,id, created_at, updated_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetAllVisibilityType :many
SELECT * FROM blogType;

-- name: GetVisibilityId :one
SELECT * FROM blogType WHERE visibilityType=$1;
