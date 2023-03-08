package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/google/uuid"
)

type SubjectRepository interface {
	//admin usecase
	Fetch(param string) (data []domain.Subject, err error)
	FetchByID(id string) (data *model.APISubjectResponse, err error)
	FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	Create(entity domain.Subject) (e *domain.Subject, err error)
	AppendGrade(subjectID uuid.UUID, GradeID uuid.UUID) (e *domain.Subject, err error)
}

type SubjectService interface {
	//admin usecase
	Fetch(param string) (data []domain.Subject, err error)
	FetchByID(id string) (data *model.APISubjectResponse, err error)
	FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	Create(request model.CreateSubjectRequest) (e *domain.Subject, err error)
	AppendGrade(request model.CreateSubjectAppendGradeRequest) (e *domain.Subject, err error)
}
