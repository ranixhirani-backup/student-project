package services
import (
	"errors"
	"studentProject/models"
	"studentProject/repository"
)

type StudentService interface {
	CreateStudent(student models.Student)(int, error)
	GetAllStudents() ([]models.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}


func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(student models.Student) (int, error){
	if student.Email == ""{
		return 0, errors.New("email is required")
	}
	//ADD MORE VALIDATIONS 
	return s.repo.Create(student)
}

func (s *studentService) GetAllStudents() ([]models.Student, error){
	return s.repo.GetAll()
}

