package model

import (
	"github.com/google/uuid"
)

type CreateAuthMemberRequest struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password *string `json:"password"`
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
}

type CreateAuthMentorRequest struct {
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	Password      *string `json:"password"`
	MentorLevelID string  `json:"mentor_level_id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	Gender        string  `json:"gender"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	Avatar        string  `json:"avatar"`
}

type APISignUpResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}
