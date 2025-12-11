package models 
import(
	"time"
)
type Entrollment struct{
	EntrollmentId int `json:"enrollment_id"`
	StudentId int `json:"student_id"`
	CourseID       int       `json:"course_id"`
    EnrollmentDate time.Time `json:"enrollment_date"`
    Status         string    `json:"status"`
}