package service

import (
	"context"
	"errors"
	"projectuasbe/app/model"
	"projectuasbe/app/repository"

	"github.com/google/uuid"
)

type LecturerService interface {
	GetProfile(ctx context.Context, lecturerID uuid.UUID) (*model.Lecturer, error)
	GetStudents(ctx context.Context, lecturerID uuid.UUID) ([]model.Student, error)
}

type lecturerServiceImpl struct {
	lecturerRepo repository.LecturerRepository
}

func NewLecturerService(repo repository.LecturerRepository) LecturerService {
	return &lecturerServiceImpl{
		lecturerRepo: repo,
	}
}

//
// GET LECTURER PROFILE
//
func (s *lecturerServiceImpl) GetProfile(ctx context.Context, lecturerID uuid.UUID) (*model.Lecturer, error) {

	if lecturerID == uuid.Nil {
		return nil, errors.New("invalid lecturer ID")
	}

	lecturer, err := s.lecturerRepo.GetByID(ctx, lecturerID)
	if err != nil {
		return nil, errors.New("lecturer not found")
	}

	return lecturer, nil
}

//
// GET ALL STUDENTS ADVISED BY THIS LECTURER
//
func (s *lecturerServiceImpl) GetStudents(ctx context.Context, lecturerID uuid.UUID) ([]model.Student, error) {

	if lecturerID == uuid.Nil {
		return nil, errors.New("invalid lecturer ID")
	}

	students, err := s.lecturerRepo.GetStudents(ctx, lecturerID)
	if err != nil {
		return nil, err
	}

	return students, nil
}
