package mentor

import (
	"encoding/json"
	"time"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"golang.org/x/crypto/bcrypt"
)

type implementAuthService struct {
	AuthRepository        usecaseMentor.AuthRepository
	MentorLevelRepository usecaseMentor.MentorLevelAdminRepository
}

// SignIn implements mentor.AuthMentorService
func (service *implementAuthService) SignIn(request request.CreateSignInMentorRequest) (accessToken string, err error) {
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

// SignUp implements mentor.AuthMentorService
func (service *implementAuthService) SignUp(request request.CreateSignUpMentorRequest) (accessToken string, err error) {
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

	mentorLevel, err := service.MentorLevelRepository.GetDataBySlug("basic")
	if err != nil {
		return "", err
	}

	var gender common.GenderType
	if request.Gender == "women" {
		gender = common.WOMEN
	} else {
		gender = common.MEN
	}

	userEntity := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Roles:    roles,
	}

	now := time.Now()
	formattedTime := now.Format("20060102150405")
	mentorEntity := domain.Mentor{
		MentorLevelID: mentorLevel.ID,
		Name:          request.Name,
		Slug:          common.MakeSlug(request.Name) + "-" + formattedTime,
		Gender:        gender,
		Phone:         request.Phone,
		Address:       request.Address,
		Avatar:        request.Avatar,
	}

	user, e := service.AuthRepository.SignUp(userEntity, mentorEntity)
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

func NewAuthService(authRepository usecaseMentor.AuthRepository, mentorLevelRepository usecaseMentor.MentorLevelAdminRepository) usecaseMentor.AuthService {
	return &implementAuthService{
		AuthRepository:        authRepository,
		MentorLevelRepository: mentorLevelRepository,
	}
}
