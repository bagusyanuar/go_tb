package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type UserRepository interface {
	Fetch() (data []model.APIUserResponse, err error)
	FetchByID(id string) (result *domain.User, err error)
	Create(user domain.User) (result *domain.User, err error)
}

type UserService interface {
	Fetch() (data []model.APIUserResponse, err error)
	FetchByID(id string) (*domain.User, error)
	Create(request model.CreateUserRequest) (user *domain.User, err error)
}
