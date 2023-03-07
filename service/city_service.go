package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type cityServiceImplementation struct {
	CityRepository usecase.CityRepository
}

func NewCityService(cityRepository usecase.CityRepository) usecase.CityService {
	return &cityServiceImplementation{
		CityRepository: cityRepository,
	}
}

// Create implements usecase.CityService
func (service *cityServiceImplementation) Create(request model.CreateCityRequest) (e *domain.City, err error) {
	provinceID, errUUIDParsing := uuid.Parse(request.ProvinceID)
	if errUUIDParsing != nil {
		return nil, errUUIDParsing
	}
	entity := domain.City{
		Name:       request.Name,
		ProvinceID: provinceID,
	}
	return service.CityRepository.Create(entity)
}

// Fetch implements usecase.CityService
func (service *cityServiceImplementation) Fetch(param string) (data []model.APICityResponse, err error) {
	return service.CityRepository.Fetch(param)
}

// FetchByID implements usecase.CityService
func (service *cityServiceImplementation) FetchByID(id string) (data *model.APICityResponse, err error) {
	return service.CityRepository.FetchByID(id)
}

// FetchBySlug implements usecase.CityService
func (service *cityServiceImplementation) FetchBySlug(slug string) (data *model.APICityResponse, err error) {
	return service.CityRepository.FetchBySlug(slug)
}
