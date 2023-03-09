package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type gradeRepositoryImplementation struct {
	Database *gorm.DB
}

func NewGradeRepository(database *gorm.DB) usecase.GradeRepository {
	return &gradeRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.GradeRepository
func (repository *gradeRepositoryImplementation) Create(entity domain.Grade) (e *domain.Grade, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.GradeRepository
func (repository *gradeRepositoryImplementation) Fetch(param string) (data []model.APIGradeResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Grade{}).Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.GradeRepository
func (repository *gradeRepositoryImplementation) FetchByID(id string) (data *model.APIGradeResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Grade{}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.GradeRepository
func (repository *gradeRepositoryImplementation) FetchBySlug(slug string) (data *model.APIGradeResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Grade{}).Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
