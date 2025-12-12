package errors
import "errors"
var ErrStudentAlreadyEnrolled = errors.New("student is already enrolled in this course")
func Wrap(msg string) error {
	return errors.New(msg)
}