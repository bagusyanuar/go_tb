package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type districtAdminRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.DistrictAdminRepository
func (repository *districtAdminRepositoryImplementation) Create(entity domain.District) (data *domain.District, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.DistrictAdminRepository
func (repository *districtAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.District{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.DistrictAdminRepository
func (repository *districtAdminRepositoryImplementation) FindAll(param string) (data []domain.District, err error) {
	if err = repository.Database.Debug().
		Preload("City").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.DistrictAdminRepository
func (repository *districtAdminRepositoryImplementation) FindByID(id string) (data *domain.District, err error) {
	if err = repository.Database.Debug().
		Preload("City").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements usecase.DistrictAdminRepository
func (repository *districtAdminRepositoryImplementation) Patch(id string, entity domain.District) (data *domain.District, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewDistrictAdminRepository(database *gorm.DB) usecase.DistrictAdminRepository {
	return &districtAdminRepositoryImplementation{
		Database: database,
	}
}
