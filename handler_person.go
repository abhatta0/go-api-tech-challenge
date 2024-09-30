package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arjunbhatta/crud/internal/database"
	"github.com/go-chi/chi/v5"
)

// handlerCreatePerson() is now a method of apiConfig struct
func (apiConf *apiConfig) handlerCreatePerson(w http.ResponseWriter, r *http.Request) {
	type parameters = struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Type      string `json:"type"`
		Age       int32  `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not parse JSON: %v", err))
		return
	}

	person, err := apiConf.DB.CreatePerson(r.Context(), database.CreatePersonParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Type:      params.Type,
		Age:       params.Age,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not create person: %v", err))
		return
	}

	respondWithJson(w, 201, databasePersonToPerson(person))
}

func (apiConf *apiConfig) handlerUpdatePerson(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	type parameters = struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Type      string `json:"type"`
		Age       int32  `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not parse JSON: %v", err))
		return
	}

	person, err := apiConf.DB.UpdatePerson(r.Context(), database.UpdatePersonParams{
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Type:        params.Type,
		Age:         params.Age,
		FirstName_2: name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not update person: %v", err))
		return
	}

	respondWithJson(w, 201, databasePersonToPerson(person))
}

func (apiConf *apiConfig) handlerGetPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := apiConf.DB.GetPersons(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not get persons: %v", err))
		return
	}

	respondWithJson(w, 201, databasePersonsToPersons(persons))
}

func (apiConf *apiConfig) handlerGetPersonByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	person, err := apiConf.DB.GetPersonByName(r.Context(), name)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not get person with name: %v, %v", name, err))
		return
	}

	respondWithJson(w, 201, databasePersonToPerson(person))
}

func (apiConf *apiConfig) handlerDeletePerson(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := apiConf.DB.DeletePerson(r.Context(), name)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR: Could not delete person with name: %v, %v", name, err))
		return
	}

	respondWithJson(w, 204, "Person deleted successfully.")
}
