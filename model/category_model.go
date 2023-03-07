package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type APICategoryResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Subjects  []withSubject  `gorm:"foreignKey:CategoryID" json:"subjects"`
}

type withSubject struct {
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Icon       *string   `json:"icon"`
	// CreatedAt  time.Time      `json:"created_at"`
	// UpdatedAt  time.Time      `json:"updated_at"`
	// DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
