package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID         uuid.UUID      `json:"id"`
	ProvinceID uuid.UUID      `json:"province_id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (city *City) BeforeCreate(tx *gorm.DB) (err error) {
	city.ID = uuid.New()
	city.CreatedAt = time.Now()
	city.UpdatedAt = time.Now()
	return
}

func (City) TableName() string {
	return "cities"
}
