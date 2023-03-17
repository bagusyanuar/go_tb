package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementMentorLevelRepository struct {
	Database *gorm.DB
}

// Create implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) Create(entity domain.MentorLevel) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&domain.MentorLevel{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) FindAll(param string) (data []domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) FindByID(id string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FindBySlug implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) FindBySlug(slug string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) Patch(id string, entity domain.MentorLevel) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewMentorLevelRepository(database *gorm.DB) usecaseAdmin.MentorLevelAdminRepository {
	return &implementMentorLevelRepository{
		Database: database,
	}
}
