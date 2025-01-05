-- name: CreateBlogs :one
INSERT INTO blogs (id, created_at, updated_at, body, title, user_id,visibility)
VALUES ($1, $2, $3, $4, $5, $6,$7)
RETURNING *;

-- name: GetUserBlogs :many
SELECT * FROM blogs WHERE user_id = $1;

-- name: GetBlog :one
SELECT * FROM blogs WHERE id=$1;

-- name: GetBlogByTilte :one
Select * FROM blogs WHERE title=$1;

-- name: UpdateUserBlog :one
UPDATE blogs
SET body = $1, title = $2, visibility = $3
WHERE id = $4 and user_id = $5
RETURNING *;

-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = $1 and user_id = $2
RETURNING *;

-- name: GetAllTypeBlogs :many
SELECT * FROM blogs WHERE visibility=$1;
