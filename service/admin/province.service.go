package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
)

type implementProvinceService struct {
	ProvinceRepository usecaseAdmin.ProvinceRepository
}

// Create implements admin.ProvinceService
func (service *implementProvinceService) Create(request request.CreateProvinceRequest) (data *domain.Province, err error) {
	entity := domain.Province{
		Name: request.Name,
		Code: request.Code,
	}
	return service.ProvinceRepository.Create(entity)
}

// Delete implements admin.ProvinceService
func (service *implementProvinceService) Delete(id string) (err error) {
	return service.ProvinceRepository.Delete(id)
}

// FindAll implements admin.ProvinceService
func (service *implementProvinceService) FindAll(param string) (data []domain.Province, err error) {
	return service.ProvinceRepository.FindAll(param)
}

// FindByID implements admin.ProvinceService
func (service *implementProvinceService) FindByID(id string) (data *domain.Province, err error) {
	return service.ProvinceRepository.FindByID(id)
}

// Patch implements admin.ProvinceService
func (service *implementProvinceService) Patch(id string, request request.CreateProvinceRequest) (data *domain.Province, err error) {
	entity := domain.Province{
		Name: request.Name,
		Code: request.Code,
	}
	return service.ProvinceRepository.Patch(id, entity)
}

func NewProvinceService(provinceRepository usecaseAdmin.ProvinceRepository) usecaseAdmin.ProvinceService {
	return &implementProvinceService{
		ProvinceRepository: provinceRepository,
	}
}
