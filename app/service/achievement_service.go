package service

import (
	"context"
	"projectuasbe/app/model"
	"projectuasbe/app/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementService struct {
	achRepo repository.AchievementRepository
	refRepo repository.AchievementReferenceRepository
}

func NewAchievementService(
	achRepo repository.AchievementRepository,
	refRepo repository.AchievementReferenceRepository,
) *AchievementService {
	return &AchievementService{
		achRepo: achRepo,
		refRepo: refRepo,
	}
}

//
// =========================
// CREATE ACHIEVEMENT
// =========================
//
func (s *AchievementService) CreateAchievement(ctx context.Context, data *model.Achievement) (string, error) {

	// Insert ke MongoDB
	res, err := s.achRepo.Create(ctx, data)
	if err != nil {
		return "", err
	}

	// Convert ID
	mongoID := res.InsertedID.(primitive.ObjectID).Hex()

	// Insert reference ke Postgres
	ref := &model.AchievementReference{
		StudentID:          data.StudentID,
		MongoAchievementID: mongoID,
		Status:             "draft",
		CreatedAt:          time.Now(),
	}

	err = s.refRepo.Create(ctx, ref)
	if err != nil {
		return "", err
	}

	return mongoID, nil
}

//
// =========================
// GET BY ID
// =========================
//
func (s *AchievementService) GetByID(ctx context.Context, id string) (*model.Achievement, error) {
	return s.achRepo.FindByID(ctx, id)
}

//
// =========================
// GET BY STUDENT
// =========================
//
func (s *AchievementService) GetByStudent(ctx context.Context, studentID string) ([]model.Achievement, error) {
	return s.achRepo.FindByStudent(ctx, studentID)
}

//
// =========================
// SUBMIT (mahasiswa)
// =========================
//
func (s *AchievementService) Submit(ctx context.Context, refID string) error {
	return s.refRepo.UpdateStatus(ctx, refID, "submitted", nil, nil)
}

//
// =========================
// VERIFY (dosen wali)
// =========================
//
func (s *AchievementService) Verify(ctx context.Context, refID string, verifiedBy string) error {
	return s.refRepo.UpdateStatus(ctx, refID, "verified", nil, &verifiedBy)
}

//
// =========================
// REJECT (dosen wali)
// =========================
//
func (s *AchievementService) Reject(ctx context.Context, refID string, note string, verifiedBy string) error {
	return s.refRepo.UpdateStatus(ctx, refID, "rejected", &note, &verifiedBy)
}

//
// =========================
// SOFT DELETE
// =========================
//
func (s *AchievementService) SoftDelete(ctx context.Context, id string) error {

	// Soft delete di Mongo
	if err := s.achRepo.SoftDelete(ctx, id); err != nil {
		return err
	}

	// Postgres soft-delete skipped because AchievementReferenceRepository
	// does not expose SoftDeleteByMongoID in the current implementation.
	// Add the appropriate repository method and restore this behavior
	// when the repository supports it.
	return nil
}
