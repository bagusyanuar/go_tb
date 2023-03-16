package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementCategoryRepository struct {
	Database *gorm.DB
}

// Create implements admin.CategoryRepository
func (repository *implementCategoryRepository) Create(entity domain.Category) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.CategoryRepository
func (repository *implementCategoryRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Category{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.CategoryRepository
func (repository *implementCategoryRepository) FindAll(param string) (data []domain.Category, err error) {
	if err = repository.Database.Debug().Preload("Subject").Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.CategoryRepository
func (repository *implementCategoryRepository) FindByID(id string) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Preload("Subject").Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.CategoryRepository
func (repository *implementCategoryRepository) Patch(id string, entity domain.Category) (data *domain.Category, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewCategoryRepository(database *gorm.DB) usecaseAdmin.CategoryRepository {
	return &implementCategoryRepository{
		Database: database,
	}
}
