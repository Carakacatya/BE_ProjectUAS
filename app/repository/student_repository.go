package repository

import (
    "context"
    "projectuasbe/app/model"
    "projectuasbe/database"
)

type StudentRepository interface {
    GetByUserID(ctx context.Context, userID string) (*model.Student, error)
    GetByID(ctx context.Context, id string) (*model.Student, error)
}

type studentRepository struct{}

func NewStudentRepository() StudentRepository {
    return &studentRepository{}
}

func (r *studentRepository) GetByUserID(ctx context.Context, userID string) (*model.Student, error) {
    db := database.Postgres

    query := `
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
        FROM students
        WHERE user_id = $1
    `

    row := db.QueryRow(ctx, query, userID)

    var s model.Student
    err := row.Scan(
        &s.ID, &s.UserID, &s.StudentID,
        &s.ProgramStudy, &s.AcademicYear,
        &s.AdvisorID, &s.CreatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &s, nil
}

func (r *studentRepository) GetByID(ctx context.Context, id string) (*model.Student, error) {
    db := database.Postgres

    query := `
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
        FROM students
        WHERE id = $1
    `

    row := db.QueryRow(ctx, query, id)

    var s model.Student
    err := row.Scan(
        &s.ID, &s.UserID, &s.StudentID,
        &s.ProgramStudy, &s.AcademicYear,
        &s.AdvisorID, &s.CreatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &s, nil
}
