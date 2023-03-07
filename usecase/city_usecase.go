package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type CityRepository interface {
	Fetch(param string) (data []model.APICityResponse, err error)
	FetchByID(id string) (data *model.APICityResponse, err error)
	FetchBySlug(slug string) (data *model.APICityResponse, err error)
	Create(entity domain.City) (e *domain.City, err error)
}

type CityService interface {
	Fetch(param string) (data []model.APICityResponse, err error)
	FetchByID(id string) (data *model.APICityResponse, err error)
	FetchBySlug(slug string) (data *model.APICityResponse, err error)
	Create(request model.CreateCityRequest) (e *domain.City, err error)
}
