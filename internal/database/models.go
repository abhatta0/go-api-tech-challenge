package database

type Course struct {
	ID   int32
	Name string
}

type Person struct {
	ID        int32
	FirstName string
	LastName  string
	Type      string
	Age       int32
}

type PersonCourse struct {
	PersonID int32
	CourseID int32
}
