package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type subjectAdminRepositoryImplementation struct {
	Database *gorm.DB
}

// AppendGrade implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) AppendGrade(id string, entity []domain.SubjectGrade) (err error) {
	tx := repository.Database.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Where("subject_id = ?", id).Delete(&domain.SubjectGrade{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

// Create implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) Create(entity domain.Subject) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Subject{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) FindAll(param string) (data []domain.Subject, err error) {
	if err = repository.Database.Debug().
		Preload("Category").
		Preload("Grades").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) FindByID(id string) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().
		Preload("Category").
		Preload("Grades").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements usecase.SubjectAdminRepository
func (repository *subjectAdminRepositoryImplementation) Patch(id string, entity domain.Subject) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewSubjectAdminRepository(database *gorm.DB) usecase.SubjectAdminRepository {
	return &subjectAdminRepositoryImplementation{
		Database: database,
	}
}
