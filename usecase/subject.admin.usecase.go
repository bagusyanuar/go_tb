package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
)

type SubjectAdminRepository interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(entity domain.Subject) (data *domain.Subject, err error)
	Patch(id string, entity domain.Subject) (data *domain.Subject, err error)
	Delete(id string) (err error)
	AppendGrade(id string, entity []domain.SubjectGrade) (err error)
}

type SubjectAdminService interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(request domain.CreateSubjectRequest) (data *domain.Subject, err error)
	Patch(id string, request domain.CreateSubjectRequest) (data *domain.Subject, err error)
	Delete(id string) (err error)
	AppendGrade(id string, request domain.CreateSubjectAppendGradeRequest) (err error)
}
