package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type CityRepository interface {
	FindAll(param string) (data []domain.City, err error)
	FindByID(id string) (data *domain.City, err error)
	Create(entity domain.City) (data *domain.City, err error)
	Patch(id string, entity domain.City) (data *domain.City, err error)
	Delete(id string) (err error)
}

type CityService interface {
	FindAll(param string) (data []domain.City, err error)
	FindByID(id string) (data *domain.City, err error)
	Create(request request.CreateCityRequest) (data *domain.City, err error)
	Patch(id string, request request.CreateCityRequest) (data *domain.City, err error)
	Delete(id string) (err error)
}
