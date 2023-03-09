package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type CategoryRepository interface {
	Fetch(param string) (data []model.APICategoryResponse, err error)
	FetchByID(id string) (data *model.APICategoryResponse, err error)
	FetchBySlug(slug string) (data *model.APICategoryResponse, err error)
	Create(entity domain.Category) (e *domain.Category, err error)
}

type CategoryService interface {
	Fetch(param string) (data []model.APICategoryResponse, err error)
	FetchByID(id string) (data *model.APICategoryResponse, err error)
	FetchBySlug(slug string) (data *model.APICategoryResponse, err error)
	Create(request model.CreateCategoryRequest) (e *domain.Category, err error)
}
