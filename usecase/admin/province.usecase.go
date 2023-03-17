package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type ProvinceRepository interface {
	FindAll(param string) (data []domain.Province, err error)
	FindByID(id string) (data *domain.Province, err error)
	Create(entity domain.Province) (data *domain.Province, err error)
	Patch(id string, entity domain.Province) (data *domain.Province, err error)
	Delete(id string) (err error)
}

type ProvinceService interface {
	FindAll(param string) (data []domain.Province, err error)
	FindByID(id string) (data *domain.Province, err error)
	Create(request request.CreateProvinceRequest) (data *domain.Province, err error)
	Patch(id string, request request.CreateProvinceRequest) (data *domain.Province, err error)
	Delete(id string) (err error)
}
