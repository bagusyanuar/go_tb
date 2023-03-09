package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type districtRepositoryImplementation struct {
	Database *gorm.DB
}

func NewDistrictRepository(database *gorm.DB) usecase.DistrictRepository {
	return &districtRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.DistrictRepository
func (repository *districtRepositoryImplementation) Create(entity domain.District) (e *domain.District, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.DistrictRepository
func (repository *districtRepositoryImplementation) Fetch(param string) (data []model.APIDistrictResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.District{}).Preload("City", func(db *gorm.DB) *gorm.DB {
		return db.Table("cities")
	}).Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.DistrictRepository
func (repository *districtRepositoryImplementation) FetchByID(id string) (data *model.APIDistrictResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.District{}).Preload("City", func(db *gorm.DB) *gorm.DB {
		return db.Table("cities")
	}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.DistrictRepository
func (repository *districtRepositoryImplementation) FetchBySlug(slug string) (data *model.APIDistrictResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.District{}).Preload("City", func(db *gorm.DB) *gorm.DB {
		return db.Table("cities")
	}).Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
