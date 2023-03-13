package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type districtAdminServiceImplementation struct {
	DistrictAdminRepository usecase.DistrictAdminRepository
}

// Create implements usecase.DistrictAdminService
func (service *districtAdminServiceImplementation) Create(request domain.CreateDistrictRequest) (data *domain.District, err error) {
	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}
	entity := domain.District{
		Name:       request.Name,
		Code:       request.Code,
		CityID: cityID,
	}
	return service.DistrictAdminRepository.Create(entity)
}

// Delete implements usecase.DistrictAdminService
func (service *districtAdminServiceImplementation) Delete(id string) (err error) {
	return service.DistrictAdminRepository.Delete(id)
}

// FindAll implements usecase.DistrictAdminService
func (service *districtAdminServiceImplementation) FindAll(param string) (data []domain.District, err error) {
	return service.DistrictAdminRepository.FindAll(param)
}

// FindByID implements usecase.DistrictAdminService
func (service *districtAdminServiceImplementation) FindByID(id string) (data *domain.District, err error) {
	return service.DistrictAdminRepository.FindByID(id)
}

// Patch implements usecase.DistrictAdminService
func (service *districtAdminServiceImplementation) Patch(id string, request domain.CreateDistrictRequest) (data *domain.District, err error) {
	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}
	entity := domain.District{
		Name:       request.Name,
		Code:       request.Code,
		CityID: cityID,
	}
	return service.DistrictAdminRepository.Patch(id, entity)
}

func NewDistrictAdminService(districtAdminRepository usecase.DistrictAdminRepository) usecase.DistrictAdminService {
	return &districtAdminServiceImplementation{
		DistrictAdminRepository: districtAdminRepository,
	}
}
