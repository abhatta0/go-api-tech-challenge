-- +goose Up
CREATE TABLE person_course
(
    person_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL,
    PRIMARY KEY (person_id, course_id),
    FOREIGN KEY (person_id) REFERENCES person (id),
    FOREIGN KEY (course_id) REFERENCES course (id)
);

-- +goose Down
DROP TABLE person_course;
