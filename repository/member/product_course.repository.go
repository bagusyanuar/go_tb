package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementProductCourseRepository struct {
	Database *gorm.DB
}

// GetData implements member.ProductCourseRepository
func (repository *implementProductCourseRepository) GetData(param string) (data []domain.ProductCourse, err error) {
	var d []domain.ProductCourse
	if err = repository.Database.Debug().Model(&domain.ProductCourse{}).Select("*", "(SELECT price FROM pricings LIMIT 1) as price").Preload("User").Find(&d).Error; err != nil {
		return data, err
	}
	return d, nil
}

func NewProductCourseRepository(database *gorm.DB) usecaseMember.ProductCourseRepository {
	return &implementProductCourseRepository{
		Database: database,
	}
}
