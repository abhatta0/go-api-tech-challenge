package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arjunbhatta/crud/internal/database"
	"github.com/go-chi/chi/v5"
)

// handlerCreateCourse() is now a method of apiConfig struct
func (apiConf *apiConfig) handlerCreateCourse(w http.ResponseWriter, r *http.Request) {
	type parameters = struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not parse JSON: %v", err))
		return
	}

	course, err := apiConf.DB.CreateCourse(r.Context(), params.Name)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not create course: %v", err))
		return
	}

	respondWithJson(w, 201, databaseCourseToCourse(course))
}

func (apiConf *apiConfig) handlerUpdateCourse(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convert the string ID to an integer
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: ID is invalid: %v", err))
		return
	}

	// Convert to int32
	id := int32(idInt)

	type parameters = struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not parse JSON: %v", err))
		return
	}

	course, err := apiConf.DB.UpdateCourse(r.Context(), database.UpdateCourseParams{
		ID:   id,
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not update course: %v", err))
		return
	}

	respondWithJson(w, 201, databaseCourseToCourse(course))
}

func (apiConf *apiConfig) handlerGetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := apiConf.DB.GetCourses(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not get courses: %v", err))
		return
	}

	respondWithJson(w, 201, databaseCoursesToCourses(courses))
}

func (apiConf *apiConfig) handlerGetCourseById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convert the string ID to an integer
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: ID is invalid: %v", err))
		return
	}

	// Convert to int32
	id := int32(idInt)

	course, err := apiConf.DB.GetCourseByID(r.Context(), id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not get course with id: %v, %v", id, err))
		return
	}

	respondWithJson(w, 201, databaseCourseToCourse(course))
}

func (apiConf *apiConfig) handlerDeleteCourse(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convert the string ID to an integer
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: ID is invalid: %v", err))
		return
	}

	// Convert to int32
	id := int32(idInt)

	err = apiConf.DB.DeleteCourse(r.Context(), id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not delete course with id: %v, %v", id, err))
		return
	}

	respondWithJson(w, 204, "Course deleted successfully.")
}
