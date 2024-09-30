-- +goose Up
CREATE TABLE course
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE course;
