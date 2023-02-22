package domain

import "context"

type User struct {
}

type UserUseCase interface {
	Fetch(ctx context.Context) ([]User, error)
}
