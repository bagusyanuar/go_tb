package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type DistrictRepository interface {
	GetData(q string) (data []domain.District, err error)
}

type DistrictService interface {
	GetData(q string) (data []model.District, err error)
}
