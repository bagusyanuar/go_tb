package usecase

import "github.com/bagusyanuar/go_tb/domain"

type CityAdminRepository interface {
	FindAll(param string) (data []domain.City, err error)
	FindByID(id string) (data *domain.City, err error)
	Create(entity domain.City) (data *domain.City, err error)
	Patch(id string, entity domain.City) (data *domain.City, err error)
	Delete(id string) (err error)
}

type CityAdminService interface {
	FindAll(param string) (data []domain.City, err error)
	FindByID(id string) (data *domain.City, err error)
	Create(request domain.CreateCityRequest) (data *domain.City, err error)
	Patch(id string, request domain.CreateCityRequest) (data *domain.City, err error)
	Delete(id string) (err error)
}
