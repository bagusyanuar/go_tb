package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementSubjectRepository struct {
	Database *gorm.DB
}

// AppendGrade implements admin.SubjectRepository
func (repository *implementSubjectRepository) AppendGrade(id string, entity []domain.SubjectGrade) (err error) {
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

// Create implements admin.SubjectRepository
func (repository *implementSubjectRepository) Create(entity domain.Subject) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.SubjectRepository
func (repository *implementSubjectRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.Subject{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.SubjectRepository
func (repository *implementSubjectRepository) FindAll(param string) (data []domain.Subject, err error) {
	if err = repository.Database.Debug().
		Preload("Category").
		Preload("Grades").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.SubjectRepository
func (repository *implementSubjectRepository) FindByID(id string) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().
		Preload("Category").
		Preload("Grades").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Patch implements admin.SubjectRepository
func (repository *implementSubjectRepository) Patch(id string, entity domain.Subject) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func NewSubjectRepository(database *gorm.DB) usecaseAdmin.SubjectRepository {
	return &implementSubjectRepository{
		Database: database,
	}
}
