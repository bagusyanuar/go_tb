package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
)

type gradeAdminServiceImplementation struct {
	GradeAdminRepository usecase.GradeAdminRepository
}

// Create implements usecase.GradeAdminService
func (service *gradeAdminServiceImplementation) Create(request domain.CreateGradeRequest) (data *domain.Grade, err error) {
	entity := domain.Grade{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.GradeAdminRepository.Create(entity)
}

// Delete implements usecase.GradeAdminService
func (service *gradeAdminServiceImplementation) Delete(id string) (err error) {
	return service.GradeAdminRepository.Delete(id)
}

// FindAll implements usecase.GradeAdminService
func (service *gradeAdminServiceImplementation) FindAll(param string) (data []domain.Grade, err error) {
	return service.GradeAdminRepository.FindAll(param)
}

// FindByID implements usecase.GradeAdminService
func (service *gradeAdminServiceImplementation) FindByID(id string) (data *domain.Grade, err error) {
	return service.GradeAdminRepository.FindByID(id)
}

// Patch implements usecase.GradeAdminService
func (service *gradeAdminServiceImplementation) Patch(id string, request domain.CreateGradeRequest) (data *domain.Grade, err error) {
	entity := domain.Grade{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.GradeAdminRepository.Patch(id, entity)
}

func NewGradeAdminService(gradeAdminRepository usecase.GradeAdminRepository) usecase.GradeAdminService {
	return &gradeAdminServiceImplementation{
		GradeAdminRepository: gradeAdminRepository,
	}
}
