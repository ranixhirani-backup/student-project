package handlers

import (
	"encoding/json"
	"net/http"
	"studentProject/models"
	"studentProject/services"
	"strconv"
    "github.com/go-chi/chi/v5"

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
	func (h *EnrollmentHandler) AcceptEnrollment(w http.ResponseWriter, r *http.Request){
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