-- name: CreateCourse :one
INSERT INTO course (name)
VALUES ($1)
RETURNING *;

-- name: GetCourses :many
SELECT * FROM course;

-- name: GetCourseByID :one
SELECT * FROM course
WHERE id = $1;

-- name: UpdateCourse :one
UPDATE course
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM course
WHERE id = $1;
