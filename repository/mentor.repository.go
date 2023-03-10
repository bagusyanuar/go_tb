package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type mentorRepositoryImplementation struct {
	Database *gorm.DB
}

func NewMentorRepository(database *gorm.DB) usecase.MentorRepository {
	return &mentorRepositoryImplementation{
		Database: database,
	}
}

// FetchByID implements usecase.MentorRepository
func (repository *mentorRepositoryImplementation) FetchByID(id string) (data *model.Mentor, err error) {
	if err = repository.Database.Debug().Model(&domain.Mentor{}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchByUserID implements usecase.MentorRepository
func (repository *mentorRepositoryImplementation) FetchByUserID(userID string) (data *model.Mentor, err error) {
	if err = repository.Database.Debug().Model(&domain.Mentor{}).Where("user_id = ?", userID).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
