package service

import (
	"encoding/json"
	"time"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImplementation struct {
	AuthRepository usecase.AuthRepository
}

func NewAuthService(authRepository usecase.AuthRepository) usecase.AuthService {
	return &authServiceImplementation{
		AuthRepository: authRepository,
	}
}

// SignUpMember implements usecase.AuthService
func (service *authServiceImplementation) SignUpMember(request model.CreateAuthMemberRequest) (accessToken string, err error) {
	roles, _ := json.Marshal([]string{"member"})
	var password *string
	if request.Password != nil {
		hashedPassword, errHashed := bcrypt.GenerateFromPassword([]byte(*request.Password), 13)
		if errHashed != nil {
			return "", errHashed
		}
		tmpPassword := string(hashedPassword)
		password = &tmpPassword
	}

	userEntity := domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: password,
		Roles:    roles,
	}

	memberEntity := domain.Member{
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
	}

	authEntity, e := service.AuthRepository.SignUpMember(userEntity, memberEntity)
	if e != nil {
		return "", e
	}
	return common.CreateJWTAccessToken(authEntity)
}

// SignUpMentor implements usecase.AuthService
func (service *authServiceImplementation) SignUpMentor(request model.CreateAuthMentorRequest) (accessToken string, err error) {
	roles, _ := json.Marshal([]string{"mentor"})
	var password *string
	if request.Password != nil {
		hashedPassword, errHashed := bcrypt.GenerateFromPassword([]byte(*request.Password), 13)
		if errHashed != nil {
			return "", errHashed
		}
		tmpPassword := string(hashedPassword)
		password = &tmpPassword
	}

	mentorLevelID, errUUIDParsing := uuid.Parse(request.MentorLevelID)
	if errUUIDParsing != nil {
		return "", errUUIDParsing
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
		MentorLevelID: mentorLevelID,
		Name:          request.Name,
		Slug:          common.MakeSlug(request.Name) + "-" + formattedTime,
		Gender:        gender,
		Phone:         request.Phone,
		Address:       request.Address,
		Avatar:        request.Avatar,
	}

	authEntity, e := service.AuthRepository.SignUpMentor(userEntity, mentorEntity)
	if e != nil {
		return "", e
	}
	return common.CreateJWTAccessToken(authEntity)
}
