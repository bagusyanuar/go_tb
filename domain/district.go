package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type District struct {
	ID        uuid.UUID      `json:"id"`
	CityID    uuid.UUID      `json:"city_id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (district *District) BeforeCreate(tx *gorm.DB) (err error) {
	district.ID = uuid.New()
	district.CreatedAt = time.Now()
	district.UpdatedAt = time.Now()
	return
}

func (District) TableName() string {
	return "districts"
}
