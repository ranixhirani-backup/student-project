package validations

import (
	"errors"
	"studentProject/models"
)

func ValidateEnrollmentPayload(e models.Enrollment) error {
	if e.StudentId == 0 {
		return errors.New("student_id is required")
	}
	if e.CourseId == 0 {
		return errors.New("course_id is required")
	}
	return nil
}