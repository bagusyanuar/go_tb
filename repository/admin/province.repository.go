package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementProvinceRepository struct {
	Database *gorm.DB
}

// Create implements admin.ProvinceRepository
func (repository *implementProvinceRepository) Create(entity domain.Province) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.ProvinceRepository
func (repository *implementProvinceRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Province{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.ProvinceRepository
func (repository *implementProvinceRepository) FindAll(param string) (data []domain.Province, err error) {
	if err = repository.Database.Debug().Preload("Cities").Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.ProvinceRepository
func (repository *implementProvinceRepository) FindByID(id string) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Preload("Cities").Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.ProvinceRepository
func (repository *implementProvinceRepository) Patch(id string, entity domain.Province) (data *domain.Province, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewProvinceRepository(database *gorm.DB) usecaseAdmin.ProvinceRepository {
	return &implementProvinceRepository{
		Database: database,
	}
}
