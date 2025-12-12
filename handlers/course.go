package handlers

import (
	"encoding/json"
	"net/http"
	"studentProject/models"
	"studentProject/services"
)
type CourseHandler struct{
	service services.CourseService
}
func NewCourseHandler(service services.CourseService) *CourseHandler{
	return &CourseHandler{service:service}
}
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request){
	var course models.Course
	if err:= json.NewDecoder(r.Body).Decode(&course); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateCourse(course)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"course_id": id,
	})
}