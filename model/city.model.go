package model

import (
	"github.com/google/uuid"
)

type CreateCityRequest struct {
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type APICityResponse struct {
	ID         uuid.UUID     `json:"id"`
	ProvinceID uuid.UUID     `json:"province_id"`
	Name       string        `json:"name"`
	Province   *withProvince `gorm:"foreignKey:ProvinceID" json:"province"`
}

type withProvince struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
