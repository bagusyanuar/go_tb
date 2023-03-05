package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"golang.org/x/crypto/bcrypt"
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

	var password *string
	if request.Password != nil {
		hashedPassword, errHashed := bcrypt.GenerateFromPassword([]byte(*request.Password), 13)
		if errHashed != nil {
			return nil, errHashed
		}
		tmpPassword := string(hashedPassword)
		password = &tmpPassword
	}

	entity := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Roles:    request.Roles,
	}
	return service.UserRepository.Create(entity)
}

// Fetch implements domain.UserService
func (service *userService) Fetch() (data []model.CreateUserResponse, err error) {
	return service.UserRepository.Fetch()
}

// FetchByID implements domain.UserService
func (*userService) FetchByID(id string) (*domain.User, error) {
	panic("unimplemented")
}
