package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type ProductCourseRepository interface {
	GetMyProductCourse(id string) (data []domain.ProductCourse, err error)
	CreateMyProductCourse(entitty domain.ProductCourse) (data *domain.ProductCourse, err error)
}

type ProductCourseService interface {
	GetMyProductCourse(id string) (data []domain.ProductCourse, err error)
	CreateMyProductCourse(userID string, request request.CreateMyProductCourseRequest) (data *domain.ProductCourse, err error)
}
