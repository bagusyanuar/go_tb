package member

import (
	"encoding/json"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"golang.org/x/crypto/bcrypt"
)

type implementAuthService struct {
	AuthRepository usecaseMember.AuthRepository
}

// SignIn implements member.AuthService
func (service *implementAuthService) SignIn(request request.CreateSignInMemberRequest) (accessToken string, err error) {
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
		Role:     "member",
	}
	return common.CreateJWTAccessToken(&jwtSign)
}

// SignUp implements member.AuthService
func (service *implementAuthService) SignUp(request request.CreateSignUpMemberRequest) (accessToken string, err error) {
	roles, _ := json.Marshal([]string{"mentor"})
	var password *string
	if request.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
		if err != nil {
			return "", err
		}
		tmp := string(hash)
		password = &tmp
	}

	userEntity := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Roles:    roles,
	}

	memberEntity := domain.Member{
		Name:    request.Name,
		Phone:   request.Phone,
		Address: request.Address,
	}

	user, e := service.AuthRepository.SignUp(userEntity, memberEntity)
	if e != nil {
		return "", e
	}

	jwtSign := common.JWTSignReturn{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     "mentor",
	}

	return common.CreateJWTAccessToken(&jwtSign)
}

func NewAuthService(authRepository usecaseMember.AuthRepository) usecaseMember.AuthService {
	return &implementAuthService{
		AuthRepository: authRepository,
	}
}
