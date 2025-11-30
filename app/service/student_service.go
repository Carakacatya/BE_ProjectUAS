package service

import (
	"projectuasbe/app/model"
	"projectuasbe/app/repository"
)

type StudentService interface {
	GetStudentByID(id string) (*model.Student, error)
	GetStudentByUserID(userID string) (*model.Student, error)
	GetStudentsByAdvisor(lecturerID string) ([]model.Student, error)
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{studentRepo: repo}
}

// =======================================
// Get student by UUID (primary key)
// =======================================
func (s *studentService) GetStudentByID(id string) (*model.Student, error) {
	return s.studentRepo.GetByID(id)
}

// =======================================
// Get student by userID (mapping User ↔ Student)
// =======================================
func (s *studentService) GetStudentByUserID(userID string) (*model.Student, error) {
	return s.studentRepo.GetByUserID(userID)
}

// =======================================
// Get all students under a lecturer advisor
// =======================================
func (s *studentService) GetStudentsByAdvisor(lecturerID string) ([]model.Student, error) {
	return s.studentRepo.GetStudentsByAdvisor(lecturerID)
}
