package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mentorLevelAdminRepositoryImplementation struct {
	Database *gorm.DB
}

// FindBySlug implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) FindBySlug(slug string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Create implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) Create(entity domain.MentorLevel) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.MentorLevel{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) FindAll(param string) (data []domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) FindByID(id string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements usecase.MentorLevelAdminRepository
func (repository *mentorLevelAdminRepositoryImplementation) Patch(id string, entity domain.MentorLevel) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewMentorLevelAdminRepository(database *gorm.DB) usecase.MentorLevelAdminRepository {
	return &mentorLevelAdminRepositoryImplementation{
		Database: database,
	}
}
