package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementProductCourseRepository struct {
	Database *gorm.DB
}

// CreateMyProductCourse implements mentor.ProductCourseRepository
func (repository *implementProductCourseRepository) CreateMyProductCourse(entitty domain.ProductCourse) (data *domain.ProductCourse, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entitty).Error; err != nil {
		return nil, err
	}
	return &entitty, nil
}

// GetMyProductCourse implements mentor.ProductCourseRepository
func (repository *implementProductCourseRepository) GetMyProductCourse(id string) (data []domain.ProductCourse, err error) {
	if err = repository.Database.Debug().
		Preload("User").
		Preload("Grade").
		Preload("Subject").
		Where("user_id = ?", id).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewProductCourseRepository(database *gorm.DB) usecaseMentor.ProductCourseRepository {
	return &implementProductCourseRepository{
		Database: database,
	}
}
