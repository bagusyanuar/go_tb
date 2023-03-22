package model

import (
	"github.com/google/uuid"
)

type City struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}

type APICityResponse struct {
	City
	Province  *CityWithProvinceScheme  `json:"province"`
	Districts []CityWithDistrictScheme `json:"districts"`
}

type CityWithProvinceScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}

type CityWithDistrictScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
