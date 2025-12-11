package main

import (
    "log"
    "net/http"

    "studentProject/db"
    "studentProject/handlers"
    "studentProject/repository"
    "studentProject/services"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    // Connect to DB
    postgresDB, err := db.NewPostgresConn(
        "localhost",
        "postgres",
        "admin",
        "student_management_db",
        "5432",
    )
    if err != nil {
        log.Fatal(err)
    }

    // Initialize layers
    studentRepo := repository.NewStudentRepository(postgresDB)
    studentService := services.NewStudentService(studentRepo)
    studentHandler := handlers.NewStudentHandler(studentService)

    // Setup Chi router
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Post("/students", studentHandler.CreateStudent)
    r.Get("/students", studentHandler.GetStudents)

    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)
}
