package models 
import(
	"time"
)
type Enrollment struct{
	EnrollmentId int `json:"enrollment_id"`
	StudentId int `json:"student_id"`
	CourseID       int       `json:"course_id"`
    EnrollmentDate time.Time `json:"enrollment_date"`
    Status         string    `json:"status"`
}