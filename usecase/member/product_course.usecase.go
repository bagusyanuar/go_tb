package member

import "github.com/bagusyanuar/go_tb/domain"

type ProductCourseRepository interface {
	GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error)
}

type ProductCourseService interface {
	GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error)
}
