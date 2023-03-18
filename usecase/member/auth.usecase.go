package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type AuthRepository interface {
	SignUp(user domain.User, member domain.Member) (data *domain.User, err error)
	SignIn(user domain.User) (data *domain.User, err error)
}

type AuthService interface {
	SignUp(request request.CreateSignUpMemberRequest) (accessToken string, err error)
	SignIn(request request.CreateSignInMemberRequest) (accessToken string, err error)
}
