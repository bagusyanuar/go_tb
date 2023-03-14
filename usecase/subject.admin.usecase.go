package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/google/uuid"
)

type SubjectAdminRepository interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(entity domain.Subject) (data *domain.Subject, err error)
	Patch(id string, entity domain.Subject) (data *domain.Subject, err error)
	Delete(id string) (err error)
	AppendGrade(id string, gradeIDS []uuid.UUID) (data *domain.Subject, err error)
}

type SubjectAdminService interface {
	FindAll(param string) (data []domain.Subject, err error)
	FindByID(id string) (data *domain.Subject, err error)
	Create(request domain.CreateSubjectRequest) (data *domain.Subject, err error)
	Patch(id string, request domain.CreateSubjectRequest) (data *domain.Subject, err error)
	Delete(id string) (err error)
}
