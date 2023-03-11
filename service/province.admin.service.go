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
	panic("unimplemented")
}

// Delete implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) Delete(id string) (err error) {
	panic("unimplemented")
}

// FindAll implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) FindAll(param string) (data []domain.Province, err error) {
	panic("unimplemented")
}

// FindByID implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) FindByID(id string) (data *domain.Province, err error) {
	panic("unimplemented")
}

// Patch implements usecase.ProvinceAdminService
func (service *provinceAdminServiceImplementation) Patch(id string, request domain.CreateProvinceRequest) (data *domain.Province, err error) {
	panic("unimplemented")
}

func NewProvinceAdminService(provinceAdminRepository usecase.ProvinceAdminRepository) usecase.ProvinceAdminService {
	return &provinceAdminServiceImplementation{
		ProvinceAdminRepository: provinceAdminRepository,
	}
}
