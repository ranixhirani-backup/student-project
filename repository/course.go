package repository
import(
	"database/sql"
	"studentProject/models"
)
type CourseRepository interface{
	Create(course models.Course) (int, error)
}
type courseRepository struct{
	DB *sql.DB
}
func NewCourseRepository(db *sql.DB) CourseRepository{
	return &courseRepository{DB: db}
}