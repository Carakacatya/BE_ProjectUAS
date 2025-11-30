package repository

import (
    "context"
    "projectuasbe/app/model"
    "projectuasbe/database"
)

type UserRepository interface {
    GetByUsername(ctx context.Context, username string) (*model.User, error)
    GetByID(ctx context.Context, id string) (*model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
    return &userRepository{}
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
    db := database.Postgres

    query := `
        SELECT id, username, email, password_hash, full_name, role_id, is_active,
               created_at, updated_at
        FROM users
        WHERE username = $1
    `

    row := db.QueryRow(ctx, query, username)

    var u model.User
    err := row.Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash,
        &u.FullName, &u.RoleID, &u.IsActive,
        &u.CreatedAt, &u.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &u, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
    db := database.Postgres

    query := `
        SELECT id, username, email, password_hash, full_name, role_id, is_active,
               created_at, updated_at
        FROM users
        WHERE id = $1
    `

    row := db.QueryRow(ctx, query, id)

    var u model.User
    err := row.Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash,
        &u.FullName, &u.RoleID, &u.IsActive,
        &u.CreatedAt, &u.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return &u, nil
}
