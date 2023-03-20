package member

import (
	"fmt"

	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementProductCourseRepository struct {
	Database *gorm.DB
}

// GetData implements member.ProductCourseRepository
func (repository *implementProductCourseRepository) GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error) {
	var d []domain.ProductCourse
	sqlPrice := repository.Database.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("pricings").Select("price").Where("grade_id = ?", gradeID).Where("city_id = ?", cityID).Where("mentor_level_id = product_courses.").Limit(1).First(&domain.Pricing{})
	})
	sqlPricing := fmt.Sprintf("(%s) as price", sqlPrice)
	if err = repository.Database.Debug().
		Model(&domain.ProductCourse{}).
		Select("*", sqlPricing).
		Where("subject_id = ?", subjectID).
		Preload("User").
		Find(&d).Error; err != nil {
		return data, err
	}
	return d, nil
}

func NewProductCourseRepository(database *gorm.DB) usecaseMember.ProductCourseRepository {
	return &implementProductCourseRepository{
		Database: database,
	}
}
