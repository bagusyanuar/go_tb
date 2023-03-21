package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type APICategoryResponse struct {
	Category
	Subjects  []CategoryWithSubjectScheme      `gorm:"foreignKey:CategoryID" json:"subjects"`
}

type CategoryWithSubjectScheme struct {
	ID         uuid.UUID      `json:"id"`
	CategoryID uuid.UUID      `json:"category_id"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	Icon       string         `json:"icon"`
}
