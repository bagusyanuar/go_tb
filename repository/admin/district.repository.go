package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementDistrictRepository struct {
	Database *gorm.DB
}

// Create implements admin.DistrictRepository
func (repository *implementDistrictRepository) Create(entity domain.District) (data *domain.District, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.DistrictRepository
func (repository *implementDistrictRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.District{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.DistrictRepository
func (repository *implementDistrictRepository) FindAll(param string) (data []domain.District, err error) {
	if err = repository.Database.Debug().
		Preload("City").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.DistrictRepository
func (repository *implementDistrictRepository) FindByID(id string) (data *domain.District, err error) {
	if err = repository.Database.Debug().
		Preload("City").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements admin.DistrictRepository
func (repository *implementDistrictRepository) Patch(id string, entity domain.District) (data *domain.District, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewDistrictRepository(database *gorm.DB) usecaseAdmin.DistrictRepository {
	return &implementDistrictRepository{
		Database: database,
	}
}
