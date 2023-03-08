package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type subjectRepositoryImplementation struct {
	Database *gorm.DB
}

func NewSubjectRepository(database *gorm.DB) usecase.SubjectRepository {
	return &subjectRepositoryImplementation{
		Database: database,
	}
}

// Create implements usecase.SubjectRepository
func (repository *subjectRepositoryImplementation) Create(entity domain.Subject) (e *domain.Subject, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.SubjectRepository
func (repository *subjectRepositoryImplementation) Fetch(param string) (data []domain.Subject, err error) {
	if err = repository.Database.Debug().Model(&domain.Subject{}).Where("name LIKE ?", "%"+param+"%").Preload("Grades").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.SubjectRepository
func (repository *subjectRepositoryImplementation) FetchByID(id string) (data *model.APISubjectResponse, err error) {
	var entity domain.Subject
	if err = repository.Database.Debug().Model(&entity).Where("id = ?", id).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Table("categories")
	}).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FetchBySlug implements usecase.SubjectRepository
func (repository *subjectRepositoryImplementation) FetchBySlug(slug string) (data *model.APISubjectResponse, err error) {
	var entity domain.Subject
	if err = repository.Database.Debug().Model(&entity).Where("slug = ?", slug).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Table("categories")
	}).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// AppendGrade implements usecase.SubjectRepository
func (repository *subjectRepositoryImplementation) AppendGrade(subjectID uuid.UUID, GradeID uuid.UUID) (e *domain.Subject, err error) {
	entity := domain.SubjectGrade{
		SubjectID: subjectID,
		GradeID:   GradeID,
	}
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
