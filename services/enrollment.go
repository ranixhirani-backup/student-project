package services

import (
	myErr "studentProject/errors"
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

	// 2. DB validation
	exists, err := s.repo.IsStudentEnrolled(enrollment.StudentId, enrollment.CourseId)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, myErr.ErrStudentAlreadyEnrolled
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

