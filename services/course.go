package services
import(
	"studentProject/models"
	"studentProject/repository"
	myErr "studentProject/errors"
)
type CourseService interface{ //interface of CourseService, any struct that would implement CreateCourse with the given signature will automatically implement this interface
	CreateCourse(course models.Course) (int, error)
}
type courseService struct{ //courseService is the concrete implementation of the CourseService interface.
							//It depends on a CourseRepository, which is injected into it.
	repo repository.CourseRepository 
}
func NewCourseService(repo repository.CourseRepository) CourseService{ //This is a constructor function.It creates a new courseService and returns it as a CourseService interface.We return a pointer because the methods have pointer receivers
		return &courseService{repo:repo}
}
func (s *courseService) CreateCourse(course models.Course) (int, error){ //implementation of the method defined in the interface
	
	exists, err := s.repo.CourseExistsByName(course.CourseName)
	if err != nil {
        // Other error, return it
        return 0, err
    }
	if exists {
		return 0, myErr.ErrCourseAlreadyExists
	}
	return s.repo.Create(course) 
}