package services

import (
	"errors"
	"studentProject/models"
	"studentProject/repository"
)

type sample struct {
	repo repository.StudentRepository
}

func (s *sample) CreateStudent(student models.Student) (int, error) {
	if student.Email == "" {
		return 0, errors.New("email is required")
	}
	//ADD MORE VALIDATIONS
	return s.repo.Create(student)
}

func (s *sample) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetAll()
}
