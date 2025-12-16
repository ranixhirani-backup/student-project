package services
import (
	"errors"
	"database/sql"
	"studentProject/models"
	"studentProject/repository"
	myErr "studentProject/errors"
)

type StudentService interface {
	CreateStudent(student models.Student)(int, error)
	GetAllStudents() ([]models.Student, error)
	GetStudent(studentId int)(models.Student, error)
	UpdateStudent(studentId int, patch models.StudentUpdate) (models.Student, error)
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

func (s *studentService) GetStudent(studentId int) (models.Student, error){
	student, err := s.repo.GetStudentById(studentId)
	if err != nil {
		if err == sql.ErrNoRows{
			return models.Student{}, myErr.ErrStudentNotFound
		}
		return models.Student{}, err
	}
	return student, nil
}

func (s *studentService)UpdateStudent(studentId int, patch models.StudentUpdate) (models.Student, error){
	student, err := s.repo.UpdateStudent(studentId, patch)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}

