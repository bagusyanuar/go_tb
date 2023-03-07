package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type provinceRepositoryImplementation struct {
	Database *gorm.DB
}

func NewProvinceRepository(database *gorm.DB) usecase.ProvinceRepository {
	return &provinceRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.ProvinceRepository
func (repository *provinceRepositoryImplementation) Create(entity domain.Province) (e *domain.Province, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.ProvinceRepository
func (repository *provinceRepositoryImplementation) Fetch(param string) (data []model.APIProvinceResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Province{}).Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.ProvinceRepository
func (repository *provinceRepositoryImplementation) FetchByID(id string) (data *model.APIProvinceResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Province{}).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.ProvinceRepository
func (repository *provinceRepositoryImplementation) FetchBySlug(slug string) (data *model.APIProvinceResponse, err error) {
	if err = repository.Database.Debug().Model(&domain.Province{}).Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
