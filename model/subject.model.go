package model

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
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type CreateSubjectRequest struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

type CreateSubjectAppendGradeRequest struct {
	SubjectID string `json:"subject_id"`
	GradeID   string `json:"grade_id"`
}

type APISubjectResponse struct {
	ID         uuid.UUID      `json:"id"`
	CategoryID uuid.UUID      `json:"category_id"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	Icon       string         `json:"icon"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Category   *Category      `gorm:"foreignKey:CategoryID" json:"category"`
	Grades     []Grade        `gorm:"many2many:subject_grades;joinForeignKey:subject_id;" json:"grades"`
}
