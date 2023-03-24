package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementCourseRequestRepository struct {
	Database *gorm.DB
}

// Send implements member.CourseRequestRepository
func (repository *implementCourseRequestRepository) Send(entity domain.CourseRequest) (data *domain.CourseRequest, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewCourseRequestRepository(database *gorm.DB) usecaseMember.CourseRequestRepository {
	return &implementCourseRequestRepository{
		Database: database,
	}
}
