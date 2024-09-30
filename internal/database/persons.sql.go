package database

import (
	"context"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO person (first_name, last_name, type, age)
VALUES ($1, $2, $3, $4)
RETURNING id, first_name, last_name, type, age
`

type CreatePersonParams struct {
	FirstName string
	LastName  string
	Type      string
	Age       int32
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson,
		arg.FirstName,
		arg.LastName,
		arg.Type,
		arg.Age,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Type,
		&i.Age,
	)
	return i, err
}

const deletePerson = `-- name: DeletePerson :exec
DELETE FROM person
WHERE CONCAT(first_name, ' ', last_name) = $1
`

func (q *Queries) DeletePerson(ctx context.Context, firstName string) error {
	_, err := q.db.ExecContext(ctx, deletePerson, firstName)
	return err
}

const getPersonByName = `-- name: GetPersonByName :one
SELECT id, first_name, last_name, type, age FROM person
WHERE CONCAT(first_name, ' ', last_name) = $1
`

func (q *Queries) GetPersonByName(ctx context.Context, firstName string) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPersonByName, firstName)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Type,
		&i.Age,
	)
	return i, err
}

const getPersons = `-- name: GetPersons :many
SELECT id, first_name, last_name, type, age FROM person
`

func (q *Queries) GetPersons(ctx context.Context) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Type,
			&i.Age,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePerson = `-- name: UpdatePerson :one
UPDATE person
SET first_name = $1, last_name = $2, type = $3, age = $4
WHERE CONCAT(first_name, ' ', last_name) = $5
RETURNING id, first_name, last_name, type, age
`

type UpdatePersonParams struct {
	FirstName   string
	LastName    string
	Type        string
	Age         int32
	FirstName_2 string
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, updatePerson,
		arg.FirstName,
		arg.LastName,
		arg.Type,
		arg.Age,
		arg.FirstName_2,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Type,
		&i.Age,
	)
	return i, err
}
