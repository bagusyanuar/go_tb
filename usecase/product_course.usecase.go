package usecase

import (
	"github.com/bagusyanuar/go_tb/model"
)

type ProductCourseRepository interface {
	//admin usecase
	Fetch(param string) (data []model.APIProductCourseResponse, err error)
	// FetchByID(id string) (data *model.APISubjectResponse, err error)
	// FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	// Create(entity domain.Subject) (e *domain.Subject, err error)
	// AppendGrade(subjectID uuid.UUID, GradeID uuid.UUID) (e *domain.Subject, err error)
}

type ProductCourseService interface {
	//admin usecase
	Fetch(param string) (data []model.APIProductCourseResponse, err error)
	// FetchByID(id string) (data *model.APISubjectResponse, err error)
	// FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	// Create(request model.CreateSubjectRequest) (e *domain.Subject, err error)
	// AppendGrade(request model.CreateSubjectAppendGradeRequest) (e *domain.Subject, err error)
}
