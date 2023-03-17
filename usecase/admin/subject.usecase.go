package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type SubjectRepository interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(entity domain.Subject) (data *domain.Subject, err error)
	Patch(id string, entity domain.Subject) (data *domain.Subject, err error)
	Delete(id string) (err error)
	AppendGrade(id string, entity []domain.SubjectGrade) (err error)
}

type SubjectService interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(request request.CreateSubjectRequest) (data *domain.Subject, err error)
	Patch(id string, request request.CreateSubjectRequest) (data *domain.Subject, err error)
	Delete(id string) (err error)
	AppendGrade(id string, request request.CreateSubjectAppendGradeRequest) (err error)
}
