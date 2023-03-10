package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type ProductCourseRepository interface {
	//admin usecase
	Create(entity domain.ProductCourse) (e *domain.ProductCourse, err error)
	Fetch(param string) (data []model.APIProductCourseResponse, err error)
	// FetchByID(id string) (data *model.APISubjectResponse, err error)
	// FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	// AppendGrade(subjectID uuid.UUID, GradeID uuid.UUID) (e *domain.Subject, err error)
}

type ProductCourseService interface {
	//admin usecase
	Create(request model.CreateProductCourseRequest) (e *domain.ProductCourse, err error)
	Fetch(param string) (data []model.APIProductCourseResponse, err error)
	// FetchByID(id string) (data *model.APISubjectResponse, err error)
	// FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	// AppendGrade(request model.CreateSubjectAppendGradeRequest) (e *domain.Subject, err error)
}
