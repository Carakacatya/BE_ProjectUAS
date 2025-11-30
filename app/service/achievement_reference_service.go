package service

import (
	"context"
	"projectuasbe/app/model"
	"projectuasbe/app/repository"
)

type AchievementReferenceService struct {
	refRepo repository.AchievementReferenceRepository
}

func NewAchievementReferenceService(refRepo repository.AchievementReferenceRepository) *AchievementReferenceService {
	return &AchievementReferenceService{
		refRepo: refRepo,
	}
}

// GET by reference ID
func (s *AchievementReferenceService) GetByID(ctx context.Context, id string) (*model.AchievementReference, error) {
	return s.refRepo.FindByID(ctx, id)
}

// GET semua reference milik mahasiswa
func (s *AchievementReferenceService) GetByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
	return s.refRepo.FindByStudent(ctx, studentID)
}

// UPDATE STATUS saja (submit / verify / reject)
func (s *AchievementReferenceService) UpdateStatus(ctx context.Context, id string, status string, note *string, verifiedBy *string) error {
	return s.refRepo.UpdateStatus(ctx, id, status, note, verifiedBy)
}

// SOFT DELETE hanya PostgreSQL
func (s *AchievementReferenceService) SoftDelete(ctx context.Context, id string) error {
	return s.refRepo.SoftDelete(ctx, id)
}
