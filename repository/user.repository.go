package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) usecase.UserRepository {
	return &userRepositoryImplementation{
		Database: database,
	}
}

type userRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.UserRepository
func (repository *userRepositoryImplementation) Create(user domain.User) (result *domain.User, err error) {
	if err = repository.Database.Debug().Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Fetch implements usecase.UserRepository
func (repository *userRepositoryImplementation) Fetch() (data []model.APIUserResponse, err error) {
	var entity []domain.User
	if err = repository.Database.Debug().Unscoped().Model(&entity).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.UserRepository
func (repository *userRepositoryImplementation) FetchByID(id string) (data *domain.User, err error) {
	if err = repository.Database.Debug().Model(&domain.User{}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
