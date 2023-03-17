package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type AuthRepository interface {
	SignUp(user domain.User, mentor domain.Mentor) (data *domain.User, err error)
	SignIn(user domain.User) (data *domain.User, err error)
}

type AuthService interface {
	SignUp(request request.CreateSignUpMentorRequest) (accessToken string, err error)
	SignIn(request request.CreateSignInMentorRequest) (accessToken string, err error)
}