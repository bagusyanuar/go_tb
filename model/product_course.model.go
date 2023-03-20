package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//base of product course model
type ProductCourse struct {
	ID          uuid.UUID      `json:"id"`
	UserID      uuid.UUID      `json:"user_id"`
	SubjectID   uuid.UUID      `json:"subject_id"`
	GradeID     uuid.UUID      `json:"grade_id"`
	Method      datatypes.JSON `json:"method"`
	Slug        string         `json:"slug"`
	IsAvailable bool           `json:"is_available"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type CreateProductCourseRequest struct {
	UserID    string `json:"user_id"`
	SubjectID string `json:"subject_id"`
	GradeID   string `json:"grade_id"`
	Method    []int  `json:"method"`
}

type APIProductCourseResponse struct {
	ProductCourse
	User    *WithUserScheme `gorm:"foreignKey:UserID;" json:"user"`
	Subject *Subject        `json:"subject"`
	Grade   *Grade          `json:"grade"`
}

type WithUserScheme struct {
	ID       uuid.UUID            `json:"id"`
	Email    string               `json:"email"`
	Username string               `json:"username"`
	Area     []WithDistrictScheme `json:"areas"`
}

type WithDistrictScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}

func (wU *WithUserScheme) TableName() string {
	return "users"
}
