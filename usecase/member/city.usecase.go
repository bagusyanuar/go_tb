package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type CityRepository interface {
	GetData(q string) (data []domain.City, err error)
}

type CityService interface {
	GetData(q string) (data []model.APICityResponse, err error)
}
