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
	Grades []SubjectWithGradeScheme `json:"grades"`
}

type SubjectWithGradeScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}
