package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	ID       uuid.UUID      `json:"id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Password *string        `json:"password"`
	Roles    datatypes.JSON `json:"roles"`
}

type CreateUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	// Password  *string        `json:"password"`
	// Roles     datatypes.JSON `json:"roles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
