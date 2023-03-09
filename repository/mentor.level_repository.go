package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type mentorLevelRepositoryImplementation struct {
	Database *gorm.DB
}

func NewMentorLevelRepository(database *gorm.DB) usecase.MentorLevelRepository {
	return &mentorLevelRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.MentorLevelRepository
func (repository *mentorLevelRepositoryImplementation) Create(entity domain.MentorLevel) (e *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.MentorLevelRepository
func (repository *mentorLevelRepositoryImplementation) Fetch(param string) (data []model.APIMentorLevelResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.MentorLevel{}).Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.MentorLevelRepository
func (repository *mentorLevelRepositoryImplementation) FetchByID(id string) (data *model.APIMentorLevelResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.MentorLevel{}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.MentorLevelRepository
func (repository *mentorLevelRepositoryImplementation) FetchBySlug(slug string) (data *model.APIMentorLevelResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.MentorLevel{}).Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
