package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	myErr "studentProject/errors"
	"studentProject/models"
	"studentProject/services"
	"studentProject/validations"
)

type EnrollmentHandler struct {
	service services.EnrollmentService
}

func NewEnrollmentHandler(service services.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{service: service}
}
func (h *EnrollmentHandler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var enrollment models.Enrollment

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validations.ValidateEnrollmentPayload(enrollment); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Service layer will handle ALL validation
	id, err := h.service.CreateEnrollment(enrollment)

	if err != nil {
		// business error
		if errors.Is(err, myErr.ErrStudentAlreadyEnrolled) {
			http.Error(w, err.Error(), http.StatusConflict) // 409
			return
		}

		// server/db error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"enrollment_id": id,
	})
}

func (h *EnrollmentHandler) AcceptEnrollment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid enrollment id", http.StatusBadRequest)
		return
	}
	err = h.service.AcceptEnrollment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Enrollment accepted successfully",
	})
}
