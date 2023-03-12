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
	Province   *Province      `gorm:"foreignKey:ProvinceID" json:"province"`
	Districts  []District     `gorm:"foreignKey:CityID" json:"districts"`
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

//request to create new city
type CreateCityRequest struct {
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
}
