package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Grade struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Subject   []Subject      `gorm:"many2many:subject_grades;" json:"subjects"`
}

func (grade *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	grade.ID = uuid.New()
	grade.CreatedAt = time.Now()
	grade.UpdatedAt = time.Now()
	return
}

func (Grade) TableName() string {
	return "grades"
}

//request to create new grade
type CreateGradeRequest struct {
	Name string `json:"name"`
}
