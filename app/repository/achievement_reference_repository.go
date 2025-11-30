package repository

import (
	"context"
	"projectuasbe/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AchievementReferenceRepository interface {
	Create(ctx context.Context, ref *model.AchievementReference) error
	UpdateStatus(ctx context.Context, id string, status string, note *string, verifiedBy *string) error
	FindByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error)
	FindByID(ctx context.Context, id string) (*model.AchievementReference, error)
	SoftDelete(ctx context.Context, id string) error
}

type achievementReferenceRepository struct {
	db *pgxpool.Pool
}

func NewAchievementReferenceRepository(db *pgxpool.Pool) AchievementReferenceRepository {
	return &achievementReferenceRepository{
		db: db,
	}
}

func (r *achievementReferenceRepository) Create(ctx context.Context, ref *model.AchievementReference) error {
	query := `
		INSERT INTO achievement_references 
		(student_id, mongo_achievement_id, status)
		VALUES ($1, $2, 'draft')
		RETURNING id;
	`

	return r.db.QueryRow(ctx, query,
		ref.StudentID, ref.MongoAchievementID,
	).Scan(&ref.ID)
}

func (r *achievementReferenceRepository) UpdateStatus(ctx context.Context, id string, status string, note *string, verifiedBy *string) error {
	query := `
		UPDATE achievement_references SET 
			status = $1,
			rejection_note = $2,
			verified_by = $3,
			updated_at = NOW()
		WHERE id = $4
	`

	_, err := r.db.Exec(ctx, query,
		status, note, verifiedBy, id,
	)

	return err
}

func (r *achievementReferenceRepository) FindByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
	query := `
		SELECT id, student_id, mongo_achievement_id, status,
			   submitted_at, verified_at, verified_by, rejection_note,
			   created_at, updated_at,
			   is_deleted, deleted_at
		FROM achievement_references
		WHERE student_id=$1 AND is_deleted=false
	`

	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.StudentID, &a.MongoAchievementID,
			&a.Status, &a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
			&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
			&a.IsDeleted, &a.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, a)
	}

	return list, nil
}

func (r *achievementReferenceRepository) FindByID(ctx context.Context, id string) (*model.AchievementReference, error) {
	query := `
		SELECT id, student_id, mongo_achievement_id, status,
			   submitted_at, verified_at, verified_by, rejection_note,
			   created_at, updated_at,
			   is_deleted, deleted_at
		FROM achievement_references
		WHERE id=$1
	`

	row := r.db.QueryRow(ctx, query, id)

	var a model.AchievementReference
	err := row.Scan(
		&a.ID, &a.StudentID, &a.MongoAchievementID,
		&a.Status, &a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
		&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
		&a.IsDeleted, &a.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *achievementReferenceRepository) SoftDelete(ctx context.Context, id string) error {
	query := `
		UPDATE achievement_references 
		SET is_deleted = true,
		    deleted_at = NOW(),
		    status = 'deleted'
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
