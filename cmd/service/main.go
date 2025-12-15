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

    courseRepo := repository.NewCourseRepository(postgresDB)
    courseService := services.NewCourseService(courseRepo)
    courseHandler := handlers.NewCourseHandler(courseService)

    enrollmentRepo := repository.NewEnrollmentRepository(postgresDB)
    enrollmentService := services.NewEnrollmentService(enrollmentRepo)
    enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentService)

    // Setup Chi router
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Post("/students", studentHandler.CreateStudent)
    r.Get("/students", studentHandler.GetStudents)
    r.Get("/students/{id}", studentHandler.GetStudentById)

    r.Post("/course", courseHandler.CreateCourse)

    r.Post("/enrollment", enrollmentHandler.CreateEnrollment)

    r.Post("/enrollments/{id}/accept", enrollmentHandler.AcceptEnrollment)
    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)
}
