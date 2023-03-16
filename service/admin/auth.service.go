package admin

import (
	"errors"
	"fmt"
	"sync"

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

func (service *authServiceImplementation) CheckConcurrent()  {
	var wg sync.WaitGroup
	errChanel := make(chan error, 2)
	jwtSign := make(chan common.JWTSignReturn, 1)
	wg.Add(2)
	go service.checkPassword(&wg, errChanel)
	go service.checkJWT(&wg, errChanel, jwtSign)
	wg.Wait()
	close(errChanel)
	close(jwtSign)
	for v := range errChanel {
		fmt.Println(v.Error())
	}
	for j := range jwtSign {
		fmt.Println(j.Email)
	}
}

func (service *authServiceImplementation) checkPassword(wg *sync.WaitGroup, err chan error)  {
	defer wg.Done()
	error := errors.New("error check password")
	err <- error
}
func (service *authServiceImplementation) checkJWT(wg *sync.WaitGroup, err chan error, jwtSign chan common.JWTSignReturn)  {
	defer wg.Done()
	error := errors.New("error create jwt")
	err <- error
	jwtSign <- common.JWTSignReturn{
		Email:    "email",
		Username: "username",
		Role:     "mentor",
	}
}


func NewAuthService(authRepository usecaseAdmin.AuthRepository) usecaseAdmin.AuthService {
	return &authServiceImplementation{
		AuthRepository: authRepository,
	}
}
