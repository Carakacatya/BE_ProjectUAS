package repository

import (
    "context"
    "projectuasbe/app/model"
    "projectuasbe/database"
)

type LecturerRepository interface {
    GetByUserID(ctx context.Context, userID string) (*model.Lecturer, error)
}

type lecturerRepository struct{}

func NewLecturerRepository() LecturerRepository {
    return &lecturerRepository{}
}

func (r *lecturerRepository) GetByUserID(ctx context.Context, userID string) (*model.Lecturer, error) {
    db := database.Postgres

    query := `
        SELECT id, user_id, lecturer_id, department, created_at
        FROM lecturers
        WHERE user_id = $1
    `

    row := db.QueryRow(ctx, query, userID)

    var l model.Lecturer
    err := row.Scan(
        &l.ID, &l.UserID, &l.LecturerID,
        &l.Department, &l.CreatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &l, nil
}
