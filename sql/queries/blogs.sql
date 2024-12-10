-- name: CreateBlogs :one
INSERT INTO blogs (id, created_at, updated_at, body, title, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserBlogs :one
SELECT * FROM blogs WHERE user_id = $1;

-- name: GetBlog :one
SELECT * FROM blogs WHERE id=$1;

-- name: UpdateUserBlog :one
UPDATE blogs
SET body = $1, title = $2, visibility = $3
WHERE id = $4
RETURNING *;

-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = $1
RETURNING *;

-- name: GetAllTypeBlogs :one
SELECT * FROM blogs WHERE visibility=$1;
