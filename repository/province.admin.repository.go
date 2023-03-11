package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type provinceAdminRepositoryImplementation struct {
	Database *gorm.DB
}

func NewProvinceAdminRepository(database *gorm.DB) usecase.ProvinceAdminRepository {
	return &provinceAdminRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.ProvinceAdminRepository
func (repository *provinceAdminRepositoryImplementation) Create(entity domain.Province) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.ProvinceAdminRepository
func (repository *provinceAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Province{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.ProvinceAdminRepository
func (repository *provinceAdminRepositoryImplementation) FindAll(param string) (data []domain.Province, err error) {
	if err = repository.Database.Debug().Preload("Cities").Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.ProvinceAdminRepository
func (repository *provinceAdminRepositoryImplementation) FindByID(id string) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Preload("Cities").Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements usecase.ProvinceAdminRepository
func (repository *provinceAdminRepositoryImplementation) Patch(id string, entity domain.Province) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
