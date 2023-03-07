package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type AuthRepository interface {
	SignUpMember(user domain.User, member domain.Member) (data *model.APISignUpResponse, err error)
}

type AuthService interface {
	SignUpMember(request model.CreateAuthMemberRequest) (accessToken string, err error)
}
