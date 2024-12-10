-- +goose Up
ALTER TABLE blogType
RENAME COLUMN blogType TO visibilityType;


-- +goose Down
ALTER TABLE blogType
RENAME COLUMN visibilityType TO blogType;
