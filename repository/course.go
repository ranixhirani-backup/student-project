package repository
import(
	"database/sql"
	"studentProject/models"
)
type CourseRepository interface{
	Create(course models.Course) (int, error) //Any struct that has a method with the exact same signature automatically implements this interface.
}
type courseRepository struct{
	DB *sql.DB //pointer to db connection
}
func NewCourseRepository(db *sql.DB) CourseRepository{
	return &courseRepository{DB: db}  //creates your repository and returns a pointer to it so other parts of your program can use it.
}

func (r *courseRepository) Create(course models.Course) (int, error){ //function of courseRepository struct type that immplements Create method which makes it of CourseRepository Interface
	var id int														 //this function takes course from model and input parameter and returns either int or error			
	query:= `
		INSERT INTO course (course_name, credits, description)
		VALUES($1, $2, $2)
		RETURNING student_id
	`
	err:= r.DB.QueryRow(query,
	course.CourseName,
	course.Credits,
	course.Description,
	).Scan(&id)
	return id,err
}