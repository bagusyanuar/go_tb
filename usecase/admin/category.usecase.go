package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type CategoryRepository interface {
	FindAll(param string) (data []domain.Category, err error)
	FindByID(id string) (data *domain.Category, err error)
	Create(entity domain.Category) (data *domain.Category, err error)
	Patch(id string, entity domain.Category) (data *domain.Category, err error)
	Delete(id string) (err error)
}

type CategoryService interface {
	FindAll(param string) (data []domain.Category, err error)
	FindByID(id string) (data *domain.Category, err error)
	Create(request request.CreateCategoryRequest) (data *domain.Category, err error)
	Patch(id string, request request.CreateCategoryRequest) (data *domain.Category, err error)
	Delete(id string) (err error)
}
