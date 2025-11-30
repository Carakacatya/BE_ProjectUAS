package repository

import (
    "context"
    "projectuasbe/app/model"
    "projectuasbe/database"
)

type AchievementReferenceRepository interface {
    Insert(ctx context.Context, ref *model.AchievementReference) error
    GetByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error)
}

type achievementReferenceRepository struct{}

func NewAchievementReferenceRepository() AchievementReferenceRepository {
    return &achievementReferenceRepository{}
}

func (r *achievementReferenceRepository) Insert(ctx context.Context, ref *model.AchievementReference) error {
    db := database.Postgres

    query := `
        INSERT INTO achievement_references 
        (id, student_id, mongo_achievement_id, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `

    _, err := db.Exec(
        ctx,
        query,
        ref.ID, ref.StudentID, ref.MongoAchievementID, ref.Status,
    )

    return err
}

func (r *achievementReferenceRepository) GetByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
    db := database.Postgres

    query := `
        SELECT id, student_id, mongo_achievement_id, status, 
               submitted_at, verified_at, verified_by, rejection_note,
               created_at, updated_at, is_deleted, deleted_at
        FROM achievement_references
        WHERE student_id = $1
    `

    rows, err := db.Query(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    refs := []model.AchievementReference{}

    for rows.Next() {
        var ref model.AchievementReference
        err := rows.Scan(
            &ref.ID, &ref.StudentID, &ref.MongoAchievementID, &ref.Status,
            &ref.SubmittedAt, &ref.VerifiedAt, &ref.VerifiedBy,
            &ref.RejectionNote, &ref.CreatedAt, &ref.UpdatedAt,
            &ref.IsDeleted, &ref.DeletedAt,
        )
        if err != nil {
            return nil, err
        }

        refs = append(refs, ref)
    }

    return refs, nil
}
