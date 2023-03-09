package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type districtServiceImplementation struct {
	DistrictRepository usecase.DistrictRepository
}

func NewDistrictService(districtRepository usecase.DistrictRepository) usecase.DistrictService {
	return &districtServiceImplementation{
		DistrictRepository: districtRepository,
	}
}

// Create implements usecase.DistrictService
func (service *districtServiceImplementation) Create(request model.CreateDistrictRequest) (e *domain.District, err error) {
	cityID, errUUIDParsing := uuid.Parse(request.CityID)
	if errUUIDParsing != nil {
		return nil, errUUIDParsing
	}
	entity := domain.District{
		Name:   request.Name,
		CityID: cityID,
	}
	return service.DistrictRepository.Create(entity)
}

// Fetch implements usecase.DistrictService
func (service *districtServiceImplementation) Fetch(param string) (data []model.APIDistrictResponse, err error) {
	return service.DistrictRepository.Fetch(param)
}

// FetchByID implements usecase.DistrictService
func (service *districtServiceImplementation) FetchByID(id string) (data *model.APIDistrictResponse, err error) {
	return service.DistrictRepository.FetchByID(id)
}

// FetchBySlug implements usecase.DistrictService
func (service *districtServiceImplementation) FetchBySlug(slug string) (data *model.APIDistrictResponse, err error) {
	return service.DistrictRepository.FetchBySlug(slug)
}
