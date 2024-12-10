-- +goose Up

CREATE TABLE blogs (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    body TEXT NOT NULL,
    title TEXT NOT NULL UNIQUE,
    visibility UUID NOT NULL REFERENCES blogType(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS blogs;