
package handlers

import (
	"net/http"
	myErr "studentProject/errors"
)

func MapDomainErrorToHTTP(err error) int {
	switch err {
	case myErr.ErrStudentNotFound:
		return http.StatusNotFound
	case myErr.ErrStudentAlreadyEnrolled:
		return http.StatusConflict
	case myErr.ErrInvalidStudentId:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
