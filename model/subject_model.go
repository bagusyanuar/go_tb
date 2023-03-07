package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateSubjectRequest struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

type APISubjectResponse struct {
	ID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CategoryID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"category_id"`
	Name       string         `gorm:"column:name" json:"name"`
	Slug       string         `gorm:"column:slug" json:"slug"`
	Icon       string         `gorm:"column:icon;type:text" json:"icon"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Category   *withCategory  `gorm:"foreignKey:CategoryID" json:"category"`
}

type withCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
	// CreatedAt  time.Time      `json:"created_at"`
	// UpdatedAt  time.Time      `json:"updated_at"`
	// DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
