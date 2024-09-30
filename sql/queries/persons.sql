-- name: CreatePerson :one
INSERT INTO person (first_name, last_name, type, age)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPersons :many
SELECT * FROM person;

-- name: GetPersonByName :one
SELECT * FROM person
WHERE CONCAT(first_name, ' ', last_name) = $1;

-- name: UpdatePerson :one
UPDATE person
SET first_name = $1, last_name = $2, type = $3, age = $4
WHERE CONCAT(first_name, ' ', last_name) = $5
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM person
WHERE CONCAT(first_name, ' ', last_name) = $1;
