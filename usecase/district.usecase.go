package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type DistrictRepository interface {
	Fetch(param string) (data []model.APIDistrictResponse, err error)
	FetchByID(id string) (data *model.APIDistrictResponse, err error)
	FetchBySlug(slug string) (data *model.APIDistrictResponse, err error)
	Create(entity domain.District) (e *domain.District, err error)
}

type DistrictService interface {
	Fetch(param string) (data []model.APIDistrictResponse, err error)
	FetchByID(id string) (data *model.APIDistrictResponse, err error)
	FetchBySlug(slug string) (data *model.APIDistrictResponse, err error)
	Create(request model.CreateDistrictRequest) (e *domain.District, err error)
}
