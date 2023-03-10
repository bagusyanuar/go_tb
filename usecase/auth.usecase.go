package usecase

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type AuthRepository interface {
	SignUpMember(user domain.User, member domain.Member) (data *common.JWTSignReturn, err error)
	SignUpMentor(user domain.User, mentor domain.Mentor) (data *common.JWTSignReturn, err error)
}

type AuthService interface {
	SignUpMember(request model.CreateAuthMemberRequest) (accessToken string, err error)
	SignUpMentor(request model.CreateAuthMentorRequest) (accessToken string, err error)
}
