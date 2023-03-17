package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementCityRepository struct {
	Database *gorm.DB
}

// Create implements admin.CityRepository
func (repository *implementCityRepository) Create(entity domain.City) (data *domain.City, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.CityRepository
func (repository *implementCityRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.City{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.CityRepository
func (repository *implementCityRepository) FindAll(param string) (data []domain.City, err error) {
	if err = repository.Database.Debug().
		Preload("Province").
		Preload("Districts").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.CityRepository
func (repository *implementCityRepository) FindByID(id string) (data *domain.City, err error) {
	if err = repository.Database.Debug().
		Preload("Province").
		Preload("Districts").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements admin.CityRepository
func (repository *implementCityRepository) Patch(id string, entity domain.City) (data *domain.City, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewCityRepository(database *gorm.DB) usecaseAdmin.CityRepository {
	return &implementCityRepository{
		Database: database,
	}
}
