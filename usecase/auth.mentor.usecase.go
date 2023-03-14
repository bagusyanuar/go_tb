package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type AuthMentorRepository interface {
	SignUp(user domain.User, mentor domain.Mentor) (data *domain.User, err error)
	SignIn(request model.CreateMentorSignInRequest) (data *domain.User, err error)
}

type AuthMentorService interface {
	SignUp(request domain.CreateSignUpMentorRequest) (accessToken string, err error)
	SignIn(request model.CreateMentorSignInRequest) (accessToken string, err error)
}
