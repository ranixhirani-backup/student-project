package services

import (
	"studentProject/errors"
	"studentProject/models"
	"studentProject/repository"
)
type EnrollmentService interface{
	CreateEnrollment(enrollment models.Enrollment) (int, error)
	AcceptEnrollment(enrollment_id int) error
}
type enrollmentService struct{
	repo repository.EnrollmentRepository
}
func NewEnrollmentService(repo repository.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{repo:repo}
}
func (s *enrollmentService) CreateEnrollment(enrollment models.Enrollment) (int, error){
	if err:= s.enrollmentValidation(enrollment); err != nil{
		return 0, err
	}
	setStatus(&enrollment)
	return s.repo.Create(enrollment)
}
func setStatus(e *models.Enrollment){
	e.Status = string(models.StatusPending)
}

func (s *enrollmentService) AcceptEnrollment(id int) error {
	return s.repo.UpdateEnrollmentStatus(id, string(models.StatusAccepted))
}

func (s *enrollmentService)enrollmentValidation(e models.Enrollment) error {
	exists, err := s.repo.IsStudentEnrolled(e.StudentId, e.CourseId)
	if err != nil{
		return err
	}
	if exists {
	 return errors.ErrAlreadyEnrolled
	}
	return nil
}