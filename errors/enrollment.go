package errors
import "errors"
var ErrStudentAlreadyEnrolled = errors.New("student is already enrolled in this course")
var ErrCourseAlreadyExists = errors.New("course already exists")
func Wrap(msg string) error {
	return errors.New(msg)
}
