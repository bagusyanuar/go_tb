package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) domain.UserRepository {
	return &userRepositoryImplementation{
		Database: database,
	}
}

type userRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements domain.UserRepository
func (repository *userRepositoryImplementation) Create(user domain.User) (result *domain.User, err error) {
	if err = repository.Database.Debug().Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Fetch implements domain.UserRepository
func (repository *userRepositoryImplementation) Fetch() (data []model.CreateUserResponse, err error) {
	var entity []domain.User
	if err = repository.Database.Debug().Unscoped().Model(&entity).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements domain.UserRepository
func (*userRepositoryImplementation) FetchByID(id string) (result *domain.User, err error) {
	panic("unimplemented")
}
