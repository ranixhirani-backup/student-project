package handlers

import (
	"encoding/json"
	"net/http"
	"errors"
	"studentProject/models"
	"studentProject/services"
	myErr "studentProject/errors"
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
		if errors.Is(err, myErr.ErrCourseAlreadyExists){
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"course_id": id,
	})
}