package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementMentorLevelRepository struct {
	Database *gorm.DB
}

// GetDataBySlug implements mentor.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) GetDataBySlug(slug string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("slug = ?", slug).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// GetData implements mentor.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) GetData(param string) (data []domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+param+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements mentor.MentorLevelAdminRepository
func (repository *implementMentorLevelRepository) GetDataByID(id string) (data *domain.MentorLevel, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewMentorLevelRepository(database *gorm.DB) usecaseMentor.MentorLevelAdminRepository {
	return &implementMentorLevelRepository{
		Database: database,
	}
}
