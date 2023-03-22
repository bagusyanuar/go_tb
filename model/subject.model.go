package model

import (
	"github.com/google/uuid"
)

type Subject struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Icon       *string   `json:"icon"`
}

type APISubjectResponse struct {
	Subject
}
