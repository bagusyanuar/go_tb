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

	if err = repository.Database.Debug().
		Model(&domain.ProductCourse{}).
		Joins("JOIN users ON users.id = product_courses.user_id").
		Joins("JOIN mentors ON users.id = mentors.user_id").
		Joins("JOIN grades ON product_courses.grade_id = grades.id").
		Joins("JOIN subjects ON product_courses.subject_id = subjects.id").
		Select("product_courses.*", sqlPricing).
		Preload("User.Mentor").
		Preload("Grade").
		Preload("Subject").
		Where("subjects.id = ?", subjectID).
		Where("grades.id = ?", gradeID).
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
