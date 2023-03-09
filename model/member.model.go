package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateMemberRequest struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

type APIMemberResponse struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID      `json:"user_id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	User      withUser       `gorm:"foreignKey:UserID" json:"user"`
}

type withUser struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
