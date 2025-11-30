package repository

import (
	"context"
	"errors"
	"projectuasbe/app/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LecturerRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Lecturer, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Lecturer, error)
	GetStudents(ctx context.Context, lecturerID uuid.UUID) ([]model.Student, error)
}

type lecturerRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewLecturerRepository(db *pgxpool.Pool) LecturerRepository {
	return &lecturerRepositoryImpl{db: db}
}

//
// GET BY ID
//
func (r *lecturerRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*model.Lecturer, error) {
	query := `
		SELECT id, user_id, lecturer_id, department
		FROM lecturers
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var lecturer model.Lecturer
	err := row.Scan(&lecturer.ID, &lecturer.UserID, &lecturer.LecturerID, &lecturer.Department)
	if err != nil {
		return nil, err
	}

	return &lecturer, nil
}

//
// GET BY USER ID
//
func (r *lecturerRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Lecturer, error) {
	query := `
		SELECT id, user_id, lecturer_id, department
		FROM lecturers
		WHERE user_id = $1
	`

	row := r.db.QueryRow(ctx, query, userID)

	var lecturer model.Lecturer
	err := row.Scan(&lecturer.ID, &lecturer.UserID, &lecturer.LecturerID, &lecturer.Department)
	if err != nil {
		return nil, err
	}

	return &lecturer, nil
}

//
// GET STUDENTS BY LECTURER
//
func (r *lecturerRepositoryImpl) GetStudents(ctx context.Context, lecturerID uuid.UUID) ([]model.Student, error) {
	query := `
		SELECT id, user_id, student_id, program_study, academic_year, advisor_id
		FROM students
		WHERE advisor_id = $1
	`

	rows, err := r.db.Query(ctx, query, lecturerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student

	for rows.Next() {
		var stu model.Student
		err := rows.Scan(
			&stu.ID,
			&stu.UserID,
			&stu.StudentID,
			&stu.ProgramStudy,
			&stu.AcademicYear,
			&stu.AdvisorID,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, stu)
	}

	if len(students) == 0 {
		return nil, errors.New("no students found")
	}

	return students, nil
}
