package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type cityAdminServiceImplementation struct {
	CityAdminRepository usecase.CityAdminRepository
}

// Create implements usecase.CityAdminService
func (service *cityAdminServiceImplementation) Create(request domain.CreateCityRequest) (data *domain.City, err error) {
	provinceID, err := uuid.Parse(request.ProvinceID)
	if err != nil {
		return nil, err
	}
	entity := domain.City{
		Name:       request.Name,
		Code:       request.Code,
		ProvinceID: provinceID,
	}
	return service.CityAdminRepository.Create(entity)
}

// Delete implements usecase.CityAdminService
func (service *cityAdminServiceImplementation) Delete(id string) (err error) {
	return service.CityAdminRepository.Delete(id)
}

// FindAll implements usecase.CityAdminService
func (service *cityAdminServiceImplementation) FindAll(param string) (data []domain.City, err error) {
	return service.CityAdminRepository.FindAll(param)
}

// FindByID implements usecase.CityAdminService
func (service *cityAdminServiceImplementation) FindByID(id string) (data *domain.City, err error) {
	return service.CityAdminRepository.FindByID(id)
}

// Patch implements usecase.CityAdminService
func (service *cityAdminServiceImplementation) Patch(id string, request domain.CreateCityRequest) (data *domain.City, err error) {
	provinceID, err := uuid.Parse(request.ProvinceID)
	if err != nil {
		return nil, err
	}
	entity := domain.City{
		Name:       request.Name,
		Code:       request.Code,
		ProvinceID: provinceID,
	}
	return service.CityAdminRepository.Patch(id, entity)
}

func NewCityAdminService(cityAdminRepository usecase.CityAdminRepository) usecase.CityAdminService {
	return &cityAdminServiceImplementation{
		CityAdminRepository: cityAdminRepository,
	}
}
