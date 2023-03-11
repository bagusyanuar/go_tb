package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
)

type categoryAdminServiceImplementation struct {
	CategoryAdminRepository usecase.CategoryAdminRepository
}

func NewCategoryAdminService(categoryAdminRepository usecase.CategoryAdminRepository) usecase.CategoryAdminService {
	return &categoryAdminServiceImplementation{
		CategoryAdminRepository: categoryAdminRepository,
	}
}

// Create implements usecase.CategoryAdminService
func (service *categoryAdminServiceImplementation) Create(request domain.CreateCategoryRequest) (data *domain.Category, err error) {
	entity := domain.Category{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.CategoryAdminRepository.Create(entity)
}

// FindAll implements usecase.CategoryAdminService
func (service *categoryAdminServiceImplementation) FindAll(param string) (data []domain.Category, err error) {
	return service.CategoryAdminRepository.FindAll(param)
}

// FindByID implements usecase.CategoryAdminService
func (service *categoryAdminServiceImplementation) FindByID(id string) (data *domain.Category, err error) {
	return service.CategoryAdminRepository.FindByID(id)
}

// Patch implements usecase.CategoryAdminService
func (service *categoryAdminServiceImplementation) Patch(id string, request domain.CreateCategoryRequest) (data *domain.Category, err error) {
	entity := domain.Category{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.CategoryAdminRepository.Patch(id, entity)
}

// Delete implements usecase.CategoryAdminService
func (service *categoryAdminServiceImplementation) Delete(id string) (err error) {
	return service.CategoryAdminRepository.Delete(id)
}
