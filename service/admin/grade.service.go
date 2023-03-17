package admin

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
)

type implementGradeService struct {
	GradeRepository usecaseAdmin.GradeRepository
}

// Create implements admin.GradeService
func (service *implementGradeService) Create(request request.CreateGradeRequest) (data *domain.Grade, err error) {
	entity := domain.Grade{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.GradeRepository.Create(entity)
}

// Delete implements admin.GradeService
func (service *implementGradeService) Delete(id string) (err error) {
	return service.GradeRepository.Delete(id)
}

// FindAll implements admin.GradeService
func (service *implementGradeService) FindAll(param string) (data []domain.Grade, err error) {
	return service.GradeRepository.FindAll(param)
}

// FindByID implements admin.GradeService
func (service *implementGradeService) FindByID(id string) (data *domain.Grade, err error) {
	return service.GradeRepository.FindByID(id)
}

// Patch implements admin.GradeService
func (service *implementGradeService) Patch(id string, request request.CreateGradeRequest) (data *domain.Grade, err error) {
	entity := domain.Grade{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.GradeRepository.Patch(id, entity)
}

func NewGradeService(gradeRepository usecaseAdmin.GradeRepository) usecaseAdmin.GradeService {
	return &implementGradeService{
		GradeRepository: gradeRepository,
	}
}
