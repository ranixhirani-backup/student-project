package repository

import (
    "database/sql"
    "studentProject/models"
)

type StudentRepository interface {
    Create(student models.Student) (int, error)
    GetAll() ([]models.Student, error)
}

type studentRepository struct {
    DB *sql.DB //contains a pointer to a SQL database connection (*sql.DB).
}

func NewStudentRepository(db *sql.DB) StudentRepository {
    return &studentRepository{DB: db}
}

func (r *studentRepository) Create(student models.Student) (int, error) {
    var id int
    query := `
        INSERT INTO student (first_name, last_name, email, gender)
        VALUES ($1, $2, $3, $4)
        RETURNING student_id
    `
    err := r.DB.QueryRow(query,
         student.FirstName,
         student.LastName,
         student.Email,
         student.Gender,
         ).Scan(&id)
    return id, err
}

func (r *studentRepository) GetAll() ([]models.Student, error){
    rows, err := r.DB.Query(`Select student_id, first_name, last_name, email, gender
     FROM student`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var students []models.Student
    for rows.Next(){
        var s models.Student
        if err := rows.Scan(&s.StudentId, &s.FirstName, &s.LastName, &s.Email, &s.Gender); err!= nil {
            return  nil, err
        }
        students = append(students, s)
    }
    return students, nil

}