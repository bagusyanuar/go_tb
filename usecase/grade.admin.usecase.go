package usecase

import "github.com/bagusyanuar/go_tb/domain"

type GradeAdminRepository interface {
	FindAll(param string) (data []domain.Grade, err error)
	FindByID(id string) (data *domain.Grade, err error)
	Create(entity domain.Grade) (data *domain.Grade, err error)
	Patch(id string, entity domain.Grade) (data *domain.Grade, err error)
	Delete(id string) (err error)
}

type GradeAdminService interface {
	FindAll(param string) (data []domain.Grade, err error)
	FindByID(id string) (data *domain.Grade, err error)
	Create(request domain.CreateGradeRequest) (data *domain.Grade, err error)
	Patch(id string, request domain.CreateGradeRequest) (data *domain.Grade, err error)
	Delete(id string) (err error)
}
