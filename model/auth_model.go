package model

import "github.com/google/uuid"

type CreateAuthMemberRequest struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password *string `json:"password"`
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
}

type APISignUpResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}
