package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryAdminRepositoryImplementation struct {
	Database *gorm.DB
}

func NewCategoryAdminRepository(database *gorm.DB) usecase.CategoryAdminRepository {
	return &categoryAdminRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.CategoryAdminRepository
func (repository *categoryAdminRepositoryImplementation) Create(entity domain.Category) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements usecase.CategoryAdminRepository
func (repository *categoryAdminRepositoryImplementation) FindAll(param string) (data []domain.Category, err error) {
	if err = repository.Database.Debug().Preload("Subject").Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.CategoryAdminRepository
func (repository *categoryAdminRepositoryImplementation) FindByID(id string) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Preload("Subject").Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements usecase.CategoryAdminRepository
func (repository *categoryAdminRepositoryImplementation) Patch(id string, entity domain.Category) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete implements usecase.CategoryAdminRepository
func (repository *categoryAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Category{}).Error; err != nil {
		return err
	}
	return nil
}
