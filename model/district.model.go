package model

import (
	"github.com/google/uuid"
)

type District struct {
	ID     uuid.UUID `json:"id"`
	CityID uuid.UUID `json:"city_id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type APIDistrictResponse struct {
	ID     uuid.UUID `json:"id"`
	CityID uuid.UUID `json:"city_id"`
	Name   string    `json:"name"`
	City   *withCity `gorm:"foreignKey:CityID" json:"city"`
}

type withCity struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
