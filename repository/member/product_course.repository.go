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
func (repository *implementProductCourseRepository) GetData(param string) (data []domain.ProductCourse, err error) {
	var d []domain.ProductCourse
	gradeID := "87c190bf-b66c-4079-9fa0-5a5189c5ba6f"
	cityID := "7d7002bb-8e08-41e7-b171-4f3cd21e9279"
	sqlPrice := repository.Database.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("pricings").Select("price").Where("grade_id = ?", gradeID).Where("city_id = ?", cityID).Limit(1).First(&domain.Pricing{})
	})
	sqlPricing := fmt.Sprintf("(%s) as price", sqlPrice)
	if err = repository.Database.Debug().
		Model(&domain.ProductCourse{}).
		Select("*", sqlPricing).
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
