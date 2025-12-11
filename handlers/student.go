package handlers

import (
	"net/http"
	"encoding/json"
	"studentProject/models"
	"studentProject/services"
)

type StudentHandler struct {
	service services.StudentService
}

func NewStudentHandler(service services.StudentService) *StudentHandler{
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request){
	var student models.Student
	    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := h.service.CreateStudent(student)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]any{
        "student_id": id,
    })
}
func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request){
   students, err := h.service.GetAllStudents()
   if err != nil{
    http.Error(w, "failed to fetch students", http.StatusInternalServerError)
    return
   }
   json.NewEncoder(w).Encode(students)
}