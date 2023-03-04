package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

type userService struct {
	UserRepository domain.UserRepository
}

// Create implements domain.UserService
func (service *userService) Create(request model.CreateUserRequest) (user *domain.User, err error) {
	password := request.Password
	entity := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Roles:    request.Roles,
	}
	user, err = service.UserRepository.Create(entity)
	return user, err
}

// Fetch implements domain.UserService
func (*userService) Fetch() ([]domain.User, error) {
	panic("unimplemented")
}

// FetchByID implements domain.UserService
func (*userService) FetchByID(id string) (*domain.User, error) {
	panic("unimplemented")
}
