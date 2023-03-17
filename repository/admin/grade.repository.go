package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementGradeRepository struct {
	Database *gorm.DB
}

// Create implements admin.GradeRepository
func (repository *implementGradeRepository) Create(entity domain.Grade) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.GradeRepository
func (repository *implementGradeRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Grade{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.GradeRepository
func (repository *implementGradeRepository) FindAll(param string) (data []domain.Grade, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.GradeRepository
func (repository *implementGradeRepository) FindByID(id string) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.GradeRepository
func (repository *implementGradeRepository) Patch(id string, entity domain.Grade) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewGradeRepository(database *gorm.DB) usecaseAdmin.GradeRepository {
	return &implementGradeRepository{
		Database: database,
	}
}
