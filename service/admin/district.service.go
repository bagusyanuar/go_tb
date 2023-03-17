package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/google/uuid"
)

type implementDistrictService struct {
	DistrictRepository usecaseAdmin.DistrictRepository
}

// Create implements admin.DistrictService
func (service *implementDistrictService) Create(request request.CreateDistrictRequest) (data *domain.District, err error) {
	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}
	entity := domain.District{
		Name:   request.Name,
		Code:   request.Code,
		CityID: cityID,
	}
	return service.DistrictRepository.Create(entity)
}

// Delete implements admin.DistrictService
func (service *implementDistrictService) Delete(id string) (err error) {
	return service.DistrictRepository.Delete(id)
}

// FindAll implements admin.DistrictService
func (service *implementDistrictService) FindAll(param string) (data []domain.District, err error) {
	return service.DistrictRepository.FindAll(param)
}

// FindByID implements admin.DistrictService
func (service *implementDistrictService) FindByID(id string) (data *domain.District, err error) {
	return service.DistrictRepository.FindByID(id)
}

// Patch implements admin.DistrictService
func (service *implementDistrictService) Patch(id string, request request.CreateDistrictRequest) (data *domain.District, err error) {
	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}
	entity := domain.District{
		Name:   request.Name,
		Code:   request.Code,
		CityID: cityID,
	}
	return service.DistrictRepository.Patch(id, entity)
}

func NewDistrictService(districtRepository usecaseAdmin.DistrictRepository) usecaseAdmin.DistrictService {
	return &implementDistrictService{
		DistrictRepository: districtRepository,
	}
}
