package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
)

type AuthMentorRepository interface {
	SignUp(user domain.User, mentor domain.Mentor) (data *domain.User, err error)
	SignIn(user domain.User) (data *domain.User, err error)
}

type AuthMentorService interface {
	SignUp(request domain.CreateSignUpMentorRequest) (accessToken string, err error)
	SignIn(request domain.CreateSignInMentorRequest) (accessToken string, err error)
}
