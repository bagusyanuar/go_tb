package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type productCourseRepositoryImplementation struct {
	Database *gorm.DB
}

func NewProductCourseRepository(database *gorm.DB) usecase.ProductCourseRepository {
	return &productCourseRepositoryImplementation{
		Database: database,
	}
}

// Fetch implements usecase.ProductCourseRepository
func (repository *productCourseRepositoryImplementation) Fetch(param string) (data []model.APIProductCourseResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.ProductCourse{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}).Preload("Subject").Preload("Grade").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Create implements usecase.ProductCourseRepository
func (repository *productCourseRepositoryImplementation) Create(entity domain.ProductCourse) (e *domain.ProductCourse, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
