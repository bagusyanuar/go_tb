package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type CategoryRepository interface {
	GetData(q string) (data []domain.Category, err error)
}

type CategoryService interface {
	GetData(q string) (data []model.APICategoryResponse, err error)
}
