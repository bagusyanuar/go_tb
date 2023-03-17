package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type GradeRepository interface {
	FindAll(param string) (data []domain.Grade, err error)
	FindByID(id string) (data *domain.Grade, err error)
	Create(entity domain.Grade) (data *domain.Grade, err error)
	Patch(id string, entity domain.Grade) (data *domain.Grade, err error)
	Delete(id string) (err error)
}

type GradeService interface {
	FindAll(param string) (data []domain.Grade, err error)
	FindByID(id string) (data *domain.Grade, err error)
	Create(request request.CreateGradeRequest) (data *domain.Grade, err error)
	Patch(id string, request request.CreateGradeRequest) (data *domain.Grade, err error)
	Delete(id string) (err error)
}
