package model

import "github.com/google/uuid"

type CreateDistrictRequest struct {
	CityID string `json:"city_id"`
	Name   string `json:"name"`
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
