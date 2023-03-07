package model

import "github.com/google/uuid"

type CreateGradeRequest struct {
	Name string `json:"name"`
}

type APIGradeResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}
