package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cityAdminRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.CityAdminRepository
func (repository *cityAdminRepositoryImplementation) Create(entity domain.City) (data *domain.City, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.CityAdminRepository
func (repository *cityAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.City{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.CityAdminRepository
func (repository *cityAdminRepositoryImplementation) FindAll(param string) (data []domain.City, err error) {
	if err = repository.Database.Debug().
		Preload("Province").
		Preload("Districts").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.CityAdminRepository
func (repository *cityAdminRepositoryImplementation) FindByID(id string) (data *domain.City, err error) {
	if err = repository.Database.Debug().
		Preload("Province").
		Preload("Districts").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements usecase.CityAdminRepository
func (repository *cityAdminRepositoryImplementation) Patch(id string, entity domain.City) (data *domain.City, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewCityAdminRepository(database *gorm.DB) usecase.CityAdminRepository {
	return &cityAdminRepositoryImplementation{
		Database: database,
	}
}
