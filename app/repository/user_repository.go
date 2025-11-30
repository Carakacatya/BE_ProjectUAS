package repository

import (
	"context"
	"projectuasbe/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Create(user *model.User) error
}

type userRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, email, full_name, password_hash, role_id, is_active, created_at, updated_at
		FROM users WHERE username=$1
	`

	row := r.DB.QueryRow(context.Background(), query, username)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.RoleID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	query := `
		SELECT id, username, email, full_name, password_hash, role_id, is_active, created_at, updated_at
		FROM users WHERE id=$1
	`

	row := r.DB.QueryRow(context.Background(), query, id)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.RoleID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash, full_name, role_id, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, true)
	`

	_, err := r.DB.Exec(context.Background(), query,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.RoleID,
	)

	return err
}
