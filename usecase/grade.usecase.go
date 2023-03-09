package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type GradeRepository interface {
	Fetch(param string) (data []model.APIGradeResponse, err error)
	FetchByID(id string) (data *model.APIGradeResponse, err error)
	FetchBySlug(slug string) (data *model.APIGradeResponse, err error)
	Create(entity domain.Grade) (e *domain.Grade, err error)
}

type GradeService interface {
	Fetch(param string) (data []model.APIGradeResponse, err error)
	FetchByID(id string) (data *model.APIGradeResponse, err error)
	FetchBySlug(slug string) (data *model.APIGradeResponse, err error)
	Create(request model.CreateGradeRequest) (e *domain.Grade, err error)
}
