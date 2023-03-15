package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
)

type authRepositoryImplementation struct {
	Database *gorm.DB
}

// SignIn implements admin.AuthRepository
func (repository *authRepositoryImplementation) SignIn(user domain.User) (data *domain.User, err error) {
	var entity domain.User
	if err = repository.Database.Debug().
		Where("JSON_SEARCH(roles, 'all', 'admin') IS NOT NULL").
		Where("email = ?", user.Email).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewAuthRepository(database *gorm.DB) usecaseAdmin.AuthRepository {
	return &authRepositoryImplementation{
		Database: database,
	}
}
