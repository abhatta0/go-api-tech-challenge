package main

import (
	"github.com/arjunbhatta/crud/internal/database"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Age       int32  `json:"age"`
}

type Course struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func databasePersonToPerson(dbPerson database.Person) Person {
	return Person{
		FirstName: dbPerson.FirstName,
		LastName:  dbPerson.LastName,
		Type:      dbPerson.Type,
		Age:       dbPerson.Age,
	}
}

func databasePersonsToPersons(dbPersons []database.Person) []Person {
	persons := []Person{}
	for _, dbPerson := range dbPersons {
		persons = append(persons, databasePersonToPerson(dbPerson))
	}

	return persons
}

func databaseCourseToCourse(dbCourse database.Course) Course {
	return Course{
		Id:   dbCourse.ID,
		Name: dbCourse.Name,
	}
}

func databaseCoursesToCourses(dbCourses []database.Course) []Course {
	courses := []Course{}
	for _, dbCourse := range dbCourses {
		courses = append(courses, databaseCourseToCourse(dbCourse))
	}

	return courses
}
