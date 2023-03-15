package admin

import "github.com/bagusyanuar/go_tb/domain"

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
	Create(request domain.CreateCategoryRequest) (data *domain.Category, err error)
	Patch(id string, request domain.CreateCategoryRequest) (data *domain.Category, err error)
	Delete(id string) (err error)
}
