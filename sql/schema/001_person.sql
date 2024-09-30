-- +goose Up
CREATE TABLE person
(
    id         SERIAL PRIMARY KEY,
    first_name TEXT                                          NOT NULL,
    last_name  TEXT                                          NOT NULL,
    type       TEXT CHECK (type IN ('professor', 'student')) NOT NULL,
    age        INTEGER                                       NOT NULL
);

-- +goose Down
DROP TABLE person;
