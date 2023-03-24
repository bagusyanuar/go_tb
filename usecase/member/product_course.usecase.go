package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type ProductCourseRepository interface {
	GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error)
}

type ProductCourseService interface {
	GetData(subjectID string, gradeID string, cityID string) (data []model.APIProductCourseResponse, err error)
}
