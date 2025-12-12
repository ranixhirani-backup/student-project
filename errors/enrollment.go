package errors
import "errors"
var (
	ErrAlreadyEnrolled = errors.New("student is already enrolled in this course")
)