-- name: CreateBlogVisibilityType :one
INSERT INTO blogType (visibilityType)
VALUES ($1)
RETURNING *;

-- name: GetAllVisibilityType :one
SELECT * FROM blogType;

-- name: GetVisibilityId :one
SELECT * FROM blogType WHERE visibilityType=$1;
