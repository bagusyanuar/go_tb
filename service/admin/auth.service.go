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
	// var wg sync.WaitGroup
	// errChanel := make(chan error, 2)
	// accessTokenChannel := make(chan string, 1)
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

	// wg.Add(2)
	// go service.checkPassword(&wg, errChanel, user, request)
	// go service.createAccessToken(&wg, errChanel, accessTokenChannel, user)
	// wg.Wait()
	// close(errChanel)
	// close(accessTokenChannel)
	// for e := range errChanel {
	// 	if e != nil {
	// 		return "", e
	// 	}
	// }
	// for jwt := range accessTokenChannel {
	// 	accessToken = jwt
	// }
	jwtSign := common.JWTSignReturn{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     "mentor",
	}
	return common.CreateJWTAccessToken(&jwtSign)
}

// func (service *authServiceImplementation) CheckConcurrent() {
// 	// var wg sync.WaitGroup
// 	// errChanel := make(chan error, 2)
// 	// jwtSign := make(chan common.JWTSignReturn, 1)
// 	// wg.Add(2)
// 	// go service.checkPassword(&wg, errChanel)
// 	// go service.checkJWT(&wg, errChanel, jwtSign)
// 	// wg.Wait()
// 	// close(errChanel)
// 	// close(jwtSign)
// 	// for v := range errChanel {
// 	// 	fmt.Println(v.Error())
// 	// }
// 	// for j := range jwtSign {
// 	// 	fmt.Println(j.Email)
// 	// }
// }

// func (service *authServiceImplementation) checkPassword(wg *sync.WaitGroup, err chan error, user *domain.User, request request.CreateAdminSignInRequest) {
// 	defer wg.Done()
// 	error := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
// 	err <- error
// }
// func (service *authServiceImplementation) createAccessToken(wg *sync.WaitGroup, err chan error, accessTokenChannel chan string, user *domain.User) {
// 	defer wg.Done()
// 	jwtSign := common.JWTSignReturn{
// 		ID:       user.ID,
// 		Email:    user.Email,
// 		Username: user.Username,
// 		Role:     "mentor",
// 	}
// 	t, e := common.CreateJWTAccessToken(&jwtSign)
// 	err <- e
// 	accessTokenChannel <- t
// }

func NewAuthService(authRepository usecaseAdmin.AuthRepository) usecaseAdmin.AuthService {
	return &authServiceImplementation{
		AuthRepository: authRepository,
	}
}
