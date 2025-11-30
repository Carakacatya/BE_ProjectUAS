package repository

import (
	"context"
	"projectuasbe/app/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository interface {
	GetByID(studentID string) (*model.Student, error)
	GetByUserID(userID string) (*model.Student, error)
	GetStudentsByAdvisor(lecturerID string) ([]model.Student, error)
}

type studentRepository struct {
	db *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) StudentRepository {
	return &studentRepository{db: db}
}

// ============================
// GET STUDENT BY STUDENT UUID
// ============================
func (r *studentRepository) GetByID(studentID string) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
		FROM students
		WHERE id = $1
	`

	var s model.Student
	err := r.db.QueryRow(ctx, query, studentID).Scan(
		&s.ID,
		&s.UserID,
		&s.StudentID,
		&s.ProgramStudy,
		&s.AcademicYear,
		&s.AdvisorID,
		&s.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

// ============================
// GET STUDENT BY USER ID
// ============================
func (r *studentRepository) GetByUserID(userID string) (*model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
		FROM students
		WHERE user_id = $1
	`

	var s model.Student
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&s.ID,
		&s.UserID,
		&s.StudentID,
		&s.ProgramStudy,
		&s.AcademicYear,
		&s.AdvisorID,
		&s.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

// =============================================
// GET ALL STUDENTS BY ACADEMIC ADVISOR (DOSEN)
// =============================================
func (r *studentRepository) GetStudentsByAdvisor(lecturerID string) ([]model.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
		FROM students
		WHERE advisor_id = $1
	`

	rows, err := r.db.Query(ctx, query, lecturerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Student

	for rows.Next() {
		var s model.Student
		err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.StudentID,
			&s.ProgramStudy,
			&s.AcademicYear,
			&s.AdvisorID,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}

	return list, nil
}
