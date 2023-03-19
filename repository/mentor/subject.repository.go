package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementSubjectRepository struct {
	Database *gorm.DB
}

// GetDataByID implements mentor.SubjectRepository
func (repository *implementSubjectRepository) GetDataByID(id string) (data *domain.Subject, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetData implements mentor.SubjectRepository
func (repository *implementSubjectRepository) GetData(param string) (data []domain.Subject, err error) {
	if err = repository.Database.Debug().
		Preload("Category").
		Preload("Grades").
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewSubjectRepository(database *gorm.DB) usecaseMentor.SubjectRepository {
	return &implementSubjectRepository{
		Database: database,
	}
}
