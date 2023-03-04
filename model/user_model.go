package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateUserRequest struct {
	ID       uuid.UUID      `json:"id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Password *string        `json:"password"`
	Roles    datatypes.JSON `json:"roles"`
}
