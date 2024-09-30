package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arjunbhatta/crud/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	dbURL := "postgresql://postgres.sxlawiwamikvexuuqtyn:arjunBhatta@65436543@aws-0-us-west-1.pooler.supabase.com:6543/postgres"

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("ERROR: Can't connect to DB")
	}

	queries := database.New(conn)
	apiConf := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Post("/persons", apiConf.handlerCreatePerson)
	v1Router.Put("/persons/{name}", apiConf.handlerUpdatePerson)
	v1Router.Get("/persons", apiConf.handlerGetPersons)
	v1Router.Get("/persons/{name}", apiConf.handlerGetPersonByName)
	v1Router.Delete("/persons/{name}", apiConf.handlerDeletePerson)

	v1Router.Post("/courses", apiConf.handlerCreateCourse)
	v1Router.Put("/courses/{id}", apiConf.handlerUpdateCourse)
	v1Router.Get("/courses", apiConf.handlerGetCourses)
	v1Router.Get("/courses/{id}", apiConf.handlerGetCourseById)
	v1Router.Delete("/courses/{id}", apiConf.handlerDeleteCourse)

	router.Mount("/api", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":8000",
	}

	fmt.Println("INFO: Server is starting on port: ", 8000)
	server_err := server.ListenAndServe()
	if server_err != nil {
		log.Fatal(server_err)
	}
}
