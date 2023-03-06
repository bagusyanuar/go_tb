package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(userRepository usecase.UserRepository) usecase.UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

type userService struct {
	UserRepository usecase.UserRepository
}

// Create implements usecase.UserService
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

// Fetch implements usecase.UserService
func (service *userService) Fetch() (data []model.APIUserResponse, err error) {
	return service.UserRepository.Fetch()
}

// FetchByID implements usecase.UserService
func (*userService) FetchByID(id string) (*domain.User, error) {
	panic("unimplemented")
}
