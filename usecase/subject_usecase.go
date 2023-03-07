package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type SubjectRepository interface {
	Fetch(param string) (data []model.APISubjectResponse, err error)
	FetchByID(id string) (data *model.APISubjectResponse, err error)
	FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	Create(entity domain.Subject) (e *domain.Subject, err error)
}

type SubjectService interface {
	Fetch(param string) (data []model.APISubjectResponse, err error)
	FetchByID(id string) (data *model.APISubjectResponse, err error)
	FetchBySlug(slug string) (data *model.APISubjectResponse, err error)
	Create(request model.CreateSubjectRequest) (e *domain.Subject, err error)
}
