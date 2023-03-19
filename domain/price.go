package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pricing struct {
	ID            uuid.UUID      `json:"id"`
	GradeID       uuid.UUID      `json:"grade_id"`
	CityID        uuid.UUID      `json:"city_id"`
	MentorLevelID uuid.UUID      `json:"mentor_level_id"`
	Method        uint           `json:"method"`
	Price         uint           `json:"price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Grade         *Grade         `gorm:"foreignKey:GradeID" json:"grade"`
	City          *City          `gorm:"foreignKey:CityID" json:"city"`
	MentorLevel   *MentorLevel   `gorm:"foreignKey:MentorLevelID" json:"mentor_level"`
}

func (pricing *Pricing) BeforeCreate(tx *gorm.DB) (err error) {
	pricing.ID = uuid.New()
	pricing.CreatedAt = time.Now()
	pricing.UpdatedAt = time.Now()
	return
}

func (Pricing) TableName() string {
	return "pricings"
}
