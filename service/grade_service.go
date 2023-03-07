package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type gradeServiceImplementation struct {
	GradeRepository usecase.GradeRepository
}

func NewGradeService(gradeRepository usecase.GradeRepository) usecase.GradeService {
	return &gradeServiceImplementation{
		GradeRepository: gradeRepository,
	}
}

// Create implements usecase.GradeService
func (service *gradeServiceImplementation) Create(request model.CreateGradeRequest) (e *domain.Grade, err error) {
	entity := domain.Grade{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.GradeRepository.Create(entity)
}

// Fetch implements usecase.GradeService
func (service *gradeServiceImplementation) Fetch(param string) (data []model.APIGradeResponse, err error) {
	return service.GradeRepository.Fetch(param)
}

// FetchByID implements usecase.GradeService
func (service *gradeServiceImplementation) FetchByID(id string) (data *model.APIGradeResponse, err error) {
	return service.GradeRepository.FetchByID(id)
}

// FetchBySlug implements usecase.GradeService
func (service *gradeServiceImplementation) FetchBySlug(slug string) (data *model.APIGradeResponse, err error) {
	return service.GradeRepository.FetchBySlug(slug)
}
