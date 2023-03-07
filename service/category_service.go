package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type categoryServiceImplementation struct {
	CategoryRepository usecase.CategoryRepository
}

// Create implements usecase.CategoryService
func (service *categoryServiceImplementation) Create(request model.CreateCategoryRequest) (e *domain.Category, err error) {
	entity := domain.Category{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.CategoryRepository.Create(entity)
}

// Fetch implements usecase.CategoryService
func (service *categoryServiceImplementation) Fetch(param string) (data []model.APICategoryResponse, err error) {
	return service.CategoryRepository.Fetch(param)
}

// FetchByID implements usecase.CategoryService
func (*categoryServiceImplementation) FetchByID(id string) (data *model.APICategoryResponse, err error) {
	panic("unimplemented")
}

// FetchBySlug implements usecase.CategoryService
func (*categoryServiceImplementation) FetchBySlug(slug string) (data *model.APICategoryResponse, err error) {
	panic("unimplemented")
}

func NewCategoryService(categoryRepository usecase.CategoryRepository) usecase.CategoryService {
	return &categoryServiceImplementation{
		CategoryRepository: categoryRepository,
	}
}
