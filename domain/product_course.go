package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductCourse struct {
	ID          uuid.UUID      `json:"id"`
	UserID      uuid.UUID      `json:"user_id"`
	SubjectID   uuid.UUID      `json:"subject_id"`
	GradeID     uuid.UUID      `json:"grade_id"`
	Method      datatypes.JSON `json:"method"`
	Slug        string         `json:"slug"`
	IsAvailable bool           `json:"is_available"`
	Price       int            `gorm:"->" json:"price"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	User        *User          `gorm:"foreignKey:UserID;" json:"user"`
	Subject     *Subject       `gorm:"foreignKey:SubjectID;" json:"subject"`
	Grade       *Grade         `gorm:"foreignKey:GradeID;" json:"grade"`
}

func (productCourse *ProductCourse) BeforeCreate(tx *gorm.DB) (err error) {
	productCourse.ID = uuid.New()
	productCourse.CreatedAt = time.Now()
	productCourse.UpdatedAt = time.Now()
	return
}

func (ProductCourse) TableName() string {
	return "product_courses"
}
