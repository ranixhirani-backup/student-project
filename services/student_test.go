 package services

// import (
// 	"database/sql"
// 	"errors"
// 	myErr "studentProject/errors"
// 	"studentProject/models"
// 	"testing"
// )

// // Mock repository
// type mockStudentRepository struct {
// 	createFunc         func(models.Student) (int, error)
// 	getAllFunc         func() ([]models.Student, error)
// 	getStudentByIdFunc func(int) (models.Student, error)
// }

// func (m *mockStudentRepository) Create(student models.Student) (int, error) {
// 	return m.createFunc(student)
// }

// func (m *mockStudentRepository) GetAll() ([]models.Student, error) {
// 	return m.getAllFunc()
// }

// func (m *mockStudentRepository) GetStudentById(studentId int) (models.Student, error) {
// 	return m.getStudentByIdFunc(studentId)
// }

// // Tests
// func TestCreateStudent(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{
// 			createFunc: func(s models.Student) (int, error) { return 1, nil },
// 		}
// 		service := NewStudentService(mockRepo)
// 		student := models.Student{Email: "test@example.com"}
// 		id, err := service.CreateStudent(student)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		if id != 1 {
// 			t.Errorf("expected id 1, got %d", id)
// 		}
// 	})

// 	t.Run("email empty", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{}
// 		service := NewStudentService(mockRepo)
// 		student := models.Student{Email: ""}
// 		_, err := service.CreateStudent(student)
// 		if err == nil || err.Error() != "email is required" {
// 			t.Errorf("expected 'email is required' error, got %v", err)
// 		}
// 	})

// 	t.Run("repo error", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{
// 			createFunc: func(s models.Student) (int, error) { return 0, errors.New("db error") },
// 		}
// 		service := NewStudentService(mockRepo)
// 		student := models.Student{Email: "test@example.com"}
// 		_, err := service.CreateStudent(student)
// 		if err == nil || err.Error() != "db error" {
// 			t.Errorf("expected 'db error', got %v", err)
// 		}
// 	})
// }

// func TestGetAllStudents(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		expected := []models.Student{{StudentId: 1, Email: "test@example.com"}}
// 		mockRepo := &mockStudentRepository{
// 			getAllFunc: func() ([]models.Student, error) { return expected, nil },
// 		}
// 		service := NewStudentService(mockRepo)
// 		students, err := service.GetAllStudents()
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		if len(students) != 1 || students[0].StudentId != 1 {
// 			t.Errorf("expected students, got %v", students)
// 		}
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{
// 			getAllFunc: func() ([]models.Student, error) { return nil, errors.New("db error") },
// 		}
// 		service := NewStudentService(mockRepo)
// 		_, err := service.GetAllStudents()
// 		if err == nil || err.Error() != "db error" {
// 			t.Errorf("expected 'db error', got %v", err)
// 		}
// 	})
// }

// func TestGetStudent(t *testing.T) {
// 	t.Run("found", func(t *testing.T) {
// 		expected := models.Student{StudentId: 1, Email: "test@example.com"}
// 		mockRepo := &mockStudentRepository{
// 			getStudentByIdFunc: func(id int) (models.Student, error) { return expected, nil },
// 		}
// 		service := NewStudentService(mockRepo)
// 		student, err := service.GetStudent(1)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		if student.StudentId != 1 {
// 			t.Errorf("expected student, got %v", student)
// 		}
// 	})

// 	t.Run("not found", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{
// 			getStudentByIdFunc: func(id int) (models.Student, error) { return models.Student{}, sql.ErrNoRows },
// 		}
// 		service := NewStudentService(mockRepo)
// 		_, err := service.GetStudent(1)
// 		if err != myErr.ErrStudentNotFound {
// 			t.Errorf("expected ErrStudentNotFound, got %v", err)
// 		}
// 	})

// 	t.Run("other error", func(t *testing.T) {
// 		mockRepo := &mockStudentRepository{
// 			getStudentByIdFunc: func(id int) (models.Student, error) { return models.Student{}, errors.New("db error") },
// 		}
// 		service := NewStudentService(mockRepo)
// 		_, err := service.GetStudent(1)
// 		if err == nil || err.Error() != "db error" {
// 			t.Errorf("expected 'db error', got %v", err)
// 		}
// 	})
// }
