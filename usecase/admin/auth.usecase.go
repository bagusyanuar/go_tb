package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type AuthRepository interface {
	SignIn(user domain.User) (data *domain.User, err error)
}

type AuthService interface {
	SignIn(request request.CreateAdminSignInRequest) (accessToken string, err error)
}
