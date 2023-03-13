package usecase

import "github.com/bagusyanuar/go_tb/domain"

type DistrictAdminRepository interface {
	FindAll(param string) (data []domain.District, err error)
	FindByID(id string) (data *domain.District, err error)
	Create(entity domain.District) (data *domain.District, err error)
	Patch(id string, entity domain.District) (data *domain.District, err error)
	Delete(id string) (err error)
}

type DistrictAdminService interface {
	FindAll(param string) (data []domain.District, err error)
	FindByID(id string) (data *domain.District, err error)
	Create(request domain.CreateDistrictRequest) (data *domain.District, err error)
	Patch(id string, request domain.CreateDistrictRequest) (data *domain.District, err error)
	Delete(id string) (err error)
}
