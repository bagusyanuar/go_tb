package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/google/uuid"
)

type implementCityService struct {
	CityService usecaseAdmin.CityRepository
}

// Create implements admin.CityService
func (service *implementCityService) Create(request request.CreateCityRequest) (data *domain.City, err error) {
	provinceID, err := uuid.Parse(request.ProvinceID)
	if err != nil {
		return nil, err
	}
	entity := domain.City{
		Name:       request.Name,
		Code:       request.Code,
		ProvinceID: provinceID,
	}
	return service.CityService.Create(entity)
}

// Delete implements admin.CityService
func (service *implementCityService) Delete(id string) (err error) {
	return service.CityService.Delete(id)
}

// FindAll implements admin.CityService
func (service *implementCityService) FindAll(param string) (data []domain.City, err error) {
	return service.CityService.FindAll(param)
}

// FindByID implements admin.CityService
func (service *implementCityService) FindByID(id string) (data *domain.City, err error) {
	return service.CityService.FindByID(id)
}

// Patch implements admin.CityService
func (service *implementCityService) Patch(id string, request request.CreateCityRequest) (data *domain.City, err error) {
	provinceID, err := uuid.Parse(request.ProvinceID)
	if err != nil {
		return nil, err
	}
	entity := domain.City{
		Name:       request.Name,
		Code:       request.Code,
		ProvinceID: provinceID,
	}
	return service.CityService.Patch(id, entity)
}

func NewCityService(cityService usecaseAdmin.CityRepository) usecaseAdmin.CityService {
	return &implementCityService{
		CityService: cityService,
	}
}
