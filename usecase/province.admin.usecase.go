package usecase

import "github.com/bagusyanuar/go_tb/domain"

type ProvinceAdminRepository interface {
	FindAll(param string) (data []domain.Province, err error)
	FindByID(id string) (data *domain.Province, err error)
	Create(entity domain.Province) (data *domain.Province, err error)
	Patch(id string, entity domain.Province) (data *domain.Province, err error)
	Delete(id string) (err error)
}

type ProvinceAdminService interface {
	FindAll(param string) (data []domain.Province, err error)
	FindByID(id string) (data *domain.Province, err error)
	Create(request domain.CreateProvinceRequest) (data *domain.Province, err error)
	Patch(id string, request domain.CreateProvinceRequest) (data *domain.Province, err error)
	Delete(id string) (err error)
}
