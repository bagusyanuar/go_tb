package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type ProvinceRepository interface {
	Fetch(param string) (data []model.APIProvinceResponse, err error)
	FetchByID(id string) (data *model.APIProvinceResponse, err error)
	FetchBySlug(slug string) (data *model.APIProvinceResponse, err error)
	Create(entity domain.Province) (e *domain.Province, err error)
}

type ProvinceService interface {
	Fetch(param string) (data []model.APIProvinceResponse, err error)
	FetchByID(id string) (data *model.APIProvinceResponse, err error)
	FetchBySlug(slug string) (data *model.APIProvinceResponse, err error)
	Create(request model.CreateProvinceRequest) (e *domain.Province, err error)
}
