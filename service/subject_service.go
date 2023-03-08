package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type subjectServiceImplementation struct {
	SubjectRepository usecase.SubjectRepository
}

func NewSubjectService(subjectRepository usecase.SubjectRepository) usecase.SubjectService {
	return &subjectServiceImplementation{
		SubjectRepository: subjectRepository,
	}
}

// Create implements usecase.SubjectService
func (service *subjectServiceImplementation) Create(request model.CreateSubjectRequest) (e *domain.Subject, err error) {
	categoryID, errUUIDParsing := uuid.Parse(request.CategoryID)
	if errUUIDParsing != nil {
		return nil, errUUIDParsing
	}
	entity := domain.Subject{
		CategoryID: categoryID,
		Name:       request.Name,
		Slug:       common.MakeSlug(request.Name),
		Icon:       nil,
	}
	return service.SubjectRepository.Create(entity)
}

// Fetch implements usecase.SubjectService
func (service *subjectServiceImplementation) Fetch(param string) (data []model.APISubjectResponse, err error) {
	return service.SubjectRepository.Fetch(param)
}

// FetchByID implements usecase.SubjectService
func (service *subjectServiceImplementation) FetchByID(id string) (data *model.APISubjectResponse, err error) {
	return service.SubjectRepository.FetchByID(id)
}

// FetchBySlug implements usecase.SubjectService
func (service *subjectServiceImplementation) FetchBySlug(slug string) (data *model.APISubjectResponse, err error) {
	return service.SubjectRepository.FetchBySlug(slug)
}

// AppendGrade implements usecase.SubjectService
func (service *subjectServiceImplementation) AppendGrade(request model.CreateSubjectAppendGradeRequest) (e *domain.Subject, err error) {
	gradeID, errUUIDParsing := uuid.Parse(request.GradeID)
	if errUUIDParsing != nil {
		return nil, errUUIDParsing
	}
	gradeEntity := domain.Grade{
		ID: gradeID,
	}
	return service.SubjectRepository.AppendGrade(request.SubjectID, gradeEntity)
}
