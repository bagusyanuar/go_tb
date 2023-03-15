package service

import (
	"encoding/json"
	"time"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/usecase"
	"golang.org/x/crypto/bcrypt"
)

type authMentorServiceImplementation struct {
	AuthMentorReposiotry       usecase.AuthMentorRepository
	MentorLevelAdminRepository usecase.MentorLevelAdminRepository
}

// SignIn implements usecase.AuthMentorService
func (service *authMentorServiceImplementation) SignIn(request domain.CreateSignInMentorRequest) (accessToken string, err error) {
	entity := domain.User{
		Email: request.Email,
	}

	user, err := service.AuthMentorReposiotry.SignIn(entity)
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

// SignUp implements usecase.AuthMentorService
func (service *authMentorServiceImplementation) SignUp(request domain.CreateSignUpMentorRequest) (accessToken string, err error) {
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

	mentorLevel, err := service.MentorLevelAdminRepository.FindBySlug("basic")
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

	user, e := service.AuthMentorReposiotry.SignUp(userEntity, mentorEntity)
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

func NewAuthMentorService(authMentorReposiotry usecase.AuthMentorRepository, mentorLevelAdminRepository usecase.MentorLevelAdminRepository) usecase.AuthMentorService {
	return &authMentorServiceImplementation{
		AuthMentorReposiotry:       authMentorReposiotry,
		MentorLevelAdminRepository: mentorLevelAdminRepository,
	}
}
