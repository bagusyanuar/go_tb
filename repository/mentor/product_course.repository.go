package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementProductCourseRepository struct {
	Database *gorm.DB
}

// CreateMyProductCourse implements mentor.ProductCourseRepository
func (repository *implementProductCourseRepository) CreateMyProductCourse(entitty domain.ProductCourse) (data *domain.ProductCourse, err error) {
	if err = repository.Database.Debug().Create(&entitty).Error; err != nil {
		return nil, err
	}
	return &entitty, nil
}

// GetMyProductCourse implements mentor.ProductCourseRepository
func (repository *implementProductCourseRepository) GetMyProductCourse(id string) (data []domain.ProductCourse, err error) {
	panic("unimplemented")
}

func NewProductCourseRepository(database *gorm.DB) usecaseMentor.ProductCourseRepository {
	return &implementProductCourseRepository{
		Database: database,
	}
}
