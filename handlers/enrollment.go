package handlers

import (
	"encoding/json"
	"net/http"
	"studentProject/models"
	"studentProject/services"
)
type EnrollmentHandler struct{
	service services.EnrollmentService
}
func NewEnrollmentHandler(service services.EnrollmentService) *EnrollmentHandler{
	return &EnrollmentHandler{service:service}
}
func (h *EnrollmentHandler) CreateEnrollment(w http.ResponseWriter, r *http.Request){
	var enrollment models.Enrollment
	if err:= json.NewDecoder(r.Body).Decode(&enrollment); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateEnrollment(enrollment)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"enrollment_id" : id,
	})
}