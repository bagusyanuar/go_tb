package model

import "github.com/google/uuid"

type CreateMentorLevelRequest struct {
	Name string `json:"name"`
}

type APIMentorLevelResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}
