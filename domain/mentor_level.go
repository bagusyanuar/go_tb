package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MentorLevel struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (mentorLevel *MentorLevel) BeforeCreate(tx *gorm.DB) (err error) {
	mentorLevel.ID = uuid.New()
	mentorLevel.CreatedAt = time.Now()
	mentorLevel.UpdatedAt = time.Now()
	return
}

func (MentorLevel) TableName() string {
	return "mentor_levels"
}
