package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type categoryRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.CategoryRepository
func (repository *categoryRepositoryImplementation) Create(entity domain.Category) (e *domain.Category, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.CategoryRepository
func (repository *categoryRepositoryImplementation) Fetch(param string) (data []model.APICategoryResponse, err error) {
	var entity []domain.Category
	if err = repository.Database.Debug().Model(&entity).Where("name LIKE ?", "%"+param+"%").Preload("Subjects", func(db *gorm.DB) *gorm.DB {
		return db.Table("subjects")
	}).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.CategoryRepository
func (*categoryRepositoryImplementation) FetchByID(id string) (data *model.APICategoryResponse, err error) {
	panic("unimplemented")
}

// FetchBySlug implements usecase.CategoryRepository
func (*categoryRepositoryImplementation) FetchBySlug(slug string) (data *model.APICategoryResponse, err error) {
	panic("unimplemented")
}

func NewCategoryRepository(database *gorm.DB) usecase.CategoryRepository {
	return &categoryRepositoryImplementation{
		Database: database,
	}
}
