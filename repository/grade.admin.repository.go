package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gradeAdminRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.GradeAdminRepository
func (repository *gradeAdminRepositoryImplementation) Create(entity domain.Grade) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.GradeAdminRepository
func (repository *gradeAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Grade{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.GradeAdminRepository
func (repository *gradeAdminRepositoryImplementation) FindAll(param string) (data []domain.Grade, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.GradeAdminRepository
func (repository *gradeAdminRepositoryImplementation) FindByID(id string) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements usecase.GradeAdminRepository
func (repository *gradeAdminRepositoryImplementation) Patch(id string, entity domain.Grade) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewGradeAdminRepository(database *gorm.DB) usecase.GradeAdminRepository {
	return &gradeAdminRepositoryImplementation{
		Database: database,
	}
}
