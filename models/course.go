package models

type Course struct {
	CourseId   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	Credits    int    `json:"credits"`
	Description string `json:"description"`
}