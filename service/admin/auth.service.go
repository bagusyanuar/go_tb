package admin

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImplementation struct {
	AuthRepository usecaseAdmin.AuthRepository
}

// SignIn implements admin.AuthService
func (service *authServiceImplementation) SignIn(request request.CreateAdminSignInRequest) (accessToken string, err error) {
	entity := domain.User{
		Email: request.Email,
	}

	user, err := service.AuthRepository.SignIn(entity)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	if err != nil {
		return "", exception.ErrorPasswordNotMatch
	}

	jwtSign := common.JWTSignReturn{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     "mentor",
	}
	return common.CreateJWTAccessToken(&jwtSign)
}

func NewAuthService(authRepository usecaseAdmin.AuthRepository) usecaseAdmin.AuthService {
	return &authServiceImplementation{
		AuthRepository: authRepository,
	}
}
