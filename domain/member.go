package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID      `json:"user_id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (member *Member) BeforeCreate(tx *gorm.DB) (err error) {
	member.ID = uuid.New()
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	return
}

func (Member) TableName() string {
	return "members"
}
