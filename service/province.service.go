package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type provinceServiceImplementation struct {
	ProvinceRepository usecase.ProvinceRepository
}

func NewProvinceRepository(provinceRepository usecase.ProvinceRepository) usecase.ProvinceService {
	return &provinceServiceImplementation{
		ProvinceRepository: provinceRepository,
	}
}

// Create implements usecase.ProvinceService
func (service *provinceServiceImplementation) Create(request model.CreateProvinceRequest) (e *domain.Province, err error) {
	entity := domain.Province{
		Name: request.Name,
	}
	return service.ProvinceRepository.Create(entity)
}

// Fetch implements usecase.ProvinceService
func (service *provinceServiceImplementation) Fetch(param string) (data []model.APIProvinceResponse, err error) {
	return service.ProvinceRepository.Fetch(param)
}

// FetchByID implements usecase.ProvinceService
func (service *provinceServiceImplementation) FetchByID(id string) (data *model.APIProvinceResponse, err error) {
	return service.ProvinceRepository.FetchByID(id)
}

// FetchBySlug implements usecase.ProvinceService
func (service *provinceServiceImplementation) FetchBySlug(slug string) (data *model.APIProvinceResponse, err error) {
	return service.ProvinceRepository.FetchBySlug(slug)
}
