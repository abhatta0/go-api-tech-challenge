package database

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO course (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateCourse(ctx context.Context, name string) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse, name)
	var i Course
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM course
WHERE id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourseByID = `-- name: GetCourseByID :one
SELECT id, name FROM course
WHERE id = $1
`

func (q *Queries) GetCourseByID(ctx context.Context, id int32) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourseByID, id)
	var i Course
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCourses = `-- name: GetCourses :many
SELECT id, name FROM course
`

func (q *Queries) GetCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const updateCourse = `-- name: UpdateCourse :one
UPDATE course
SET name = $2
WHERE id = $1
RETURNING id, name
`

type UpdateCourseParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourse, arg.ID, arg.Name)
	var i Course
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
