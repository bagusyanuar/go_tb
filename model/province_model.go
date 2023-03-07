package model

import "github.com/google/uuid"

type CreateProvinceRequest struct {
	Name string `json:"name"`
}

type APIProvinceResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
