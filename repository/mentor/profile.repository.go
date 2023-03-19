package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	mentorUsecase "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type profileRepositoryImplementation struct {
	Database *gorm.DB
}

// GetMyslug implements mentor.ProfileRepository
func (repository *profileRepositoryImplementation) GetMyslug(id string) (slug string, err error) {
	var data domain.Mentor
	if err = repository.Database.Debug().Where("user_id = ?", id).First(&data).Error; err != nil {
		return "", err
	}
	return data.Slug, nil
}

// GetProfile implements mentor.ProfileRepository
func (repository *profileRepositoryImplementation) GetProfile(id string) (data *domain.User, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Preload("Mentor").First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProfileRepository(database *gorm.DB) mentorUsecase.ProfileRepository {
	return &profileRepositoryImplementation{
		Database: database,
	}
}
