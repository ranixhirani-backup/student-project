package repository

import (
    "database/sql"
    "studentProject/models"
    "strconv"
)

type StudentRepository interface {
    Create(student models.Student) (int, error)
    GetAll() ([]models.Student, error)
    GetStudentById(studentId int) (models.Student, error)
    UpdateStudent(studentId int, patch models.StudentUpdate) (models.Student, error)

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

func (r *studentRepository) GetStudentById(studentId int) (models.Student, error){
    query := `Select student_id, first_name, last_name, email, gender
     FROM student
     WHERE student_id = $1
     `
    var student models.Student
    err:= r.DB.QueryRow(query, studentId).Scan(
        &student.StudentId,
        &student.FirstName,
        &student.LastName,
        &student.Email,
        &student.Gender,
    )
    return student, err

}

func (r *studentRepository) UpdateStudent(studentId int, patch models.StudentUpdate) (models.Student, error) {

    query := "UPDATE student SET "
    args := []interface{}{}
    idx := 1

    if patch.FirstName != nil {
        query += "first_name = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.FirstName)
        idx++
    }
    if patch.LastName != nil {
        query += "last_name = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.LastName)
        idx++
    }
    if patch.Email != nil {
        query += "email = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.Email)
        idx++
    }
    if patch.Gender != nil {
        query += "gender = $" + strconv.Itoa(idx) + ", "
        args = append(args, *patch.Gender)
        idx++
    }

    if len(args) == 0 {
        // nothing to update, just return current student
        return r.GetStudentById(studentId)
    }

    // remove trailing comma and space
    query = query[:len(query)-2]

    // add WHERE clause
    query += " WHERE student_id = $" + strconv.Itoa(idx)
    args = append(args, studentId)

    // execute update
    _, err := r.DB.Exec(query, args...)
    if err != nil {
        return models.Student{}, err
    }

    // return the updated student
    return r.GetStudentById(studentId)
}
