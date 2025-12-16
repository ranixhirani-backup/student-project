package models

import (
	"time"
)

type Student struct {
	StudentId int       `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

// create table  Student (student_id INT, first_name var_char(225), last_name var_char(225), email var_char(225), gender var_char(225), created_at DateTime)

type StudentUpdate struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Gender    *string `json:"gender"`

	// 	*string lets Go tell the difference between
	// “field not sent” and “field sent with empty value”.
}
