package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
)

type provinceAdminServiceImplementation struct {
	ProvinceAdminRepository usecase.ProvinceAdminRepository
}

// Create implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) Create(request domain.CreateProvinceRequest) (data *domain.Province, err error) {
	entity := domain.Province{
		Name: request.Name,
		Code: request.Code,
	}
	return service.ProvinceAdminRepository.Create(entity)
}

// Delete implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) Delete(id string) (err error) {
	return service.ProvinceAdminRepository.Delete(id)
}

// FindAll implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) FindAll(param string) (data []domain.Province, err error) {
	return service.ProvinceAdminRepository.FindAll(param)
}

// FindByID implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) FindByID(id string) (data *domain.Province, err error) {
	return service.ProvinceAdminRepository.FindByID(id)
}

// Patch implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) Patch(id string, request domain.CreateProvinceRequest) (data *domain.Province, err error) {
	entity := domain.Province{
		Name: request.Name,
		Code: request.Code,
	}
	return service.ProvinceAdminRepository.Patch(id, entity)
}

func NewProvinceAdminService(provinceAdminRepository usecase.ProvinceAdminRepository) usecase.ProvinceAdminService {
	return &provinceAdminServiceImplementation{
		ProvinceAdminRepository: provinceAdminRepository,
	}
}
