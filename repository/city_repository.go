package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type cityRepositoryImplementation struct {
	Database *gorm.DB
}

func NewCityRepository(database *gorm.DB) usecase.CityRepository {
	return &cityRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.CityRepository
func (repository *cityRepositoryImplementation) Create(entity domain.City) (e *domain.City, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.CityRepository
func (repository *cityRepositoryImplementation) Fetch(param string) (data []model.APICityResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.City{}).Preload("Province", func(db *gorm.DB) *gorm.DB {
		return db.Table("provinces")
	}).Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.CityRepository
func (repository *cityRepositoryImplementation) FetchByID(id string) (data *model.APICityResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.City{}).Preload("Province", func(db *gorm.DB) *gorm.DB {
		return db.Table("provinces")
	}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.CityRepository
func (repository *cityRepositoryImplementation) FetchBySlug(slug string) (data *model.APICityResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.City{}).Preload("Province", func(db *gorm.DB) *gorm.DB {
		return db.Table("provinces")
	}).Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
