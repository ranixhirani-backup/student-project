package repository
import(
	"database/sql"
	"studentProject/models"
)
type CourseRepository interface{
	Create(course models.Course) (int, error) //Any struct that has a method with the exact same signature automatically implements this interface.
	CourseExists(courseId int) (bool, error)
}
type courseRepository struct{
	DB *sql.DB //pointer to db connection
}
func NewCourseRepository(db *sql.DB) CourseRepository{
	return &courseRepository{DB: db}  //creates your repository and returns a pointer to it so other parts of your program can use it.
}

func (r *courseRepository) Create(course models.Course) (int, error){ //function of courseRepository struct type that implements Create method which makes it of CourseRepository Interface
	var id int														 //this function takes course from model and input parameter and returns either int or error			
	query:= `
		INSERT INTO course (course_name, credits, description)
		VALUES($1, $2, $3)
		RETURNING course_id
	`
	err:= r.DB.QueryRow(query,
	course.CourseName,
	course.Credits,
	course.Description,
	).Scan(&id)
	return id,err
}

func(r *courseRepository) CourseExists(courseId int) (bool, error){
	query := `
		SELECT 1
		FROM course
		WHERE course_id = $1 LIMIT 1
	`
	var count int
	err := r.DB.QueryRow(query, courseId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}