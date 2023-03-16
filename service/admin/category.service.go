package admin

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
)

type implementationCategoryService struct {
	CategoryRepository usecaseAdmin.CategoryRepository
}

// Create implements admin.CategoryService
func (service *implementationCategoryService) Create(request request.CreateCategoryRequest) (data *domain.Category, err error) {
	entity := domain.Category{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.CategoryRepository.Create(entity)
}

// Delete implements admin.CategoryService
func (service *implementationCategoryService) Delete(id string) (err error) {
	return service.CategoryRepository.Delete(id)
}

// FindAll implements admin.CategoryService
func (service *implementationCategoryService) FindAll(param string) (data []domain.Category, err error) {
	return service.CategoryRepository.FindAll(param)
}

// FindByID implements admin.CategoryService
func (service *implementationCategoryService) FindByID(id string) (data *domain.Category, err error) {
	return service.CategoryRepository.FindByID(id)
}

// Patch implements admin.CategoryService
func (service *implementationCategoryService) Patch(id string, request request.CreateCategoryRequest) (data *domain.Category, err error) {
	entity := domain.Category{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.CategoryRepository.Patch(id, entity)
}

func NewCategoryService(categoryRepository usecaseAdmin.CategoryRepository) usecaseAdmin.CategoryService {
	return &implementationCategoryService{
		CategoryRepository: categoryRepository,
	}
}
