package repository

import (
	"database/sql"
	"studentProject/models"
)
type EnrollmentRepository interface{
	Create(enrollment models.Enrollment) (int, error)
	UpdateEnrollmentStatus(enrollment_id int, status string) (error)
	IsStudentEnrolled(studentID, courseId int) (bool, error)
}
type enrollmentRepository struct{
	DB *sql.DB
}
func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository{
	return &enrollmentRepository{DB: db}
}
func (r *enrollmentRepository) Create(enrollment models.Enrollment) (int, error){
	var id int
	query := `
		INSERT INTO enrollment (student_id, course_id, status)
		VALUES($1, $2, $3)
		RETURNING enrollment_id
	`
	err:= r.DB.QueryRow(query,
	enrollment.StudentId,
	enrollment.CourseId,
	enrollment.Status,
	).Scan(&id)
	return id, err
}
func (r *enrollmentRepository) UpdateEnrollmentStatus(enrollment_id int, status string) (error) {
	query := `
		UPDATE enrollment 
		SET status = $1
		WHERE enrollment_id = $2
	`
	_, err:= r.DB.Exec(query, status, enrollment_id)
	return err
}

func(r *enrollmentRepository) IsStudentEnrolled(studentId, courseId int) (bool, error){
	query := `
		SELECT COUNT(*)
		FROM enrollment
		WHERE student_id = $1 AND course_id = $2
	`
	var count int
	err := r.DB.QueryRow(query, studentId, courseId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}