package model

import "github.com/google/uuid"

type Lecturer struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	LecturerID string    `json:"lecturer_id"`
	Department string    `json:"department"`
}
