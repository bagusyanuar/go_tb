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

// Check implements member.ProductCourseRepository
func (repository *implementProductCourseRepository) Check() (data []domain.ProductCourse, err error) {
	if err = repository.Database.Debug().
		Model(&domain.ProductCourse{}).
		Preload("User").
		Preload("User.Areas").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetData implements member.ProductCourseRepository
func (repository *implementProductCourseRepository) GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error) {
	// var d []domain.ProductCourse
	sqlPrice := repository.Database.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("pricings").
			Select("price").
			Where("grade_id = ?", gradeID).
			Where("city_id = ?", cityID).
			Where("mentor_level_id = `mentors`.`mentor_level_id`").
			Limit(1).First(&domain.Pricing{})
	})
	sqlPricing := fmt.Sprintf("(%s) as price", sqlPrice)
	havingPricing := fmt.Sprintf("(%s) > 0", sqlPrice)
	// if err = repository.Database.Debug().
	// 	Model(&domain.ProductCourse{}).
	// 	Select("*", sqlPricing).
	// 	Where("subject_id = ?", subjectID).
	// 	Preload("User").
	// 	Find(&d).Error; err != nil {
	// 	return data, err
	// }
	// return d, nil
	if err = repository.Database.Debug().
		Model(&domain.ProductCourse{}).
		Joins("JOIN users ON users.id = product_courses.user_id").
		Joins("JOIN mentors ON users.id = mentors.user_id").
		Joins("JOIN grades ON product_courses.grade_id = grades.id").
		Select("product_courses.*", sqlPricing).
		Preload("User.Mentor").
		Preload("Grade").
		Having(havingPricing).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewProductCourseRepository(database *gorm.DB) usecaseMember.ProductCourseRepository {
	return &implementProductCourseRepository{
		Database: database,
	}
}
