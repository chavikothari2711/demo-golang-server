-- +goose Up
CREATE TABLE blogType(
    id UUID PRIMARY KEY,
    blogType TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE blogType;