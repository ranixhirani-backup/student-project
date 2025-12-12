package repository
import(
	"database/sql"
	"studentProject/models"
)
type EnrollmentRepository interface{
	Create(enrollment models.Enrollment) (int, error)
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
		INSERT INTO enrollment (student_id, course_id)
		VALUES($1, $2)
		RETURNING enrollment_id
	`
	err:= r.DB.QueryRow(query,
	enrollment.StudentId,
	enrollment.CourseID,
	).Scan(&id)
	return id, err
}