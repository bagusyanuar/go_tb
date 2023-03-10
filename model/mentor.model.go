package model

import (
	"time"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//base mentor
type Mentor struct {
	ID            uuid.UUID         `json:"id"`
	UserID        uuid.UUID         `json:"user_id"`
	MentorLevelID uuid.UUID         `json:"mentor_level_id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	Gender        common.GenderType `json:"gender"`
	Phone         string            `json:"phone"`
	Address       string            `json:"address"`
	Avatar        string            `json:"avatar"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at"`
}
