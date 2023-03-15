package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
)

type ProfileRepository interface {
	GetProfile(id string) (data *domain.User, err error)
}

type CategoryService interface {
	GetProfile(id string) (data *domain.User, err error)
}
