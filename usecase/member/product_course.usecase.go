package member

import "github.com/bagusyanuar/go_tb/domain"

type ProductCourseRepository interface {
	GetData(param string) (data []domain.ProductCourse, err error)
}

type ProductCourseService interface {
	GetData(param string) (data []domain.ProductCourse, err error)
}
