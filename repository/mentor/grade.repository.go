package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementGradeRepository struct {
	Database *gorm.DB
}

// GetData implements mentor.GradeRepository
func (repository *implementGradeRepository) GetData(param string) (data []domain.Grade, err error) {
	panic("unimplemented")
}

// GetDataByID implements mentor.GradeRepository
func (repository *implementGradeRepository) GetDataByID(id string) (data *domain.Grade, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// GetDataBySlug implements mentor.GradeRepository
func (repository *implementGradeRepository) GetDataBySlug(slug string) (data *domain.Grade, err error) {
	panic("unimplemented")
}

func NewGradeRepository(database *gorm.DB) usecaseMentor.GradeRepository {
	return &implementGradeRepository{
		Database: database,
	}
}
