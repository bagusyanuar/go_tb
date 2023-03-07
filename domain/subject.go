package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID         uuid.UUID      `json:"id"`
	CategoryID uuid.UUID      `json:"category_id"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	Icon       *string        `json:"icon"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

func (subject *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	subject.ID = uuid.New()
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()
	return
}

func (Subject) TableName() string {
	return "subjects"
}
