package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type DistrictRepository interface {
	FindAll(param string) (data []domain.District, err error)
	FindByID(id string) (data *domain.District, err error)
	Create(entity domain.District) (data *domain.District, err error)
	Patch(id string, entity domain.District) (data *domain.District, err error)
	Delete(id string) (err error)
}

type DistrictService interface {
	FindAll(param string) (data []domain.District, err error)
	FindByID(id string) (data *domain.District, err error)
	Create(request request.CreateDistrictRequest) (data *domain.District, err error)
	Patch(id string, request request.CreateDistrictRequest) (data *domain.District, err error)
	Delete(id string) (err error)
}
