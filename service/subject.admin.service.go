package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/google/uuid"
)

type subjectAdminServiceImplementation struct {
	SubjectAdminRepository usecase.SubjectAdminRepository
}

// AppendGrade implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) AppendGrade(id string, request domain.CreateSubjectAppendGradeRequest) (err error) {
	var gradeIDS []uuid.UUID
	subjectId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	for _, gradeID := range request.GradeIDS {
		gradeIDUUID, err := uuid.Parse(gradeID)
		if err != nil {
			return err
		}
		gradeIDS = append(gradeIDS, gradeIDUUID)
	}

	entity := []domain.SubjectGrade{}
	for _, gradeID := range gradeIDS {
		entity = append(entity, domain.SubjectGrade{
			SubjectID: subjectId,
			GradeID:   gradeID,
		})
	}
	return service.SubjectAdminRepository.AppendGrade(id, entity)
}

// Create implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) Create(request domain.CreateSubjectRequest) (data *domain.Subject, err error) {
	categoryID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return nil, err
	}
	entity := domain.Subject{
		Name:       request.Name,
		Slug:       common.MakeSlug(request.Name),
		CategoryID: categoryID,
	}
	return service.SubjectAdminRepository.Create(entity)
}

// Delete implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) Delete(id string) (err error) {
	return service.SubjectAdminRepository.Delete(id)
}

// FindAll implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) FindAll(param string) (data []domain.Subject, err error) {
	return service.SubjectAdminRepository.FindAll(param)
}

// FindByID implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) FindByID(id string) (data *domain.Subject, err error) {
	return service.SubjectAdminRepository.FindByID(id)
}

// Patch implements usecase.SubjectAdminService
func (service *subjectAdminServiceImplementation) Patch(id string, request domain.CreateSubjectRequest) (data *domain.Subject, err error) {
	categoryID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return nil, err
	}
	entity := domain.Subject{
		Name:       request.Name,
		Slug:       common.MakeSlug(request.Name),
		CategoryID: categoryID,
	}
	return service.SubjectAdminRepository.Patch(id, entity)
}

func NewSubjectAdminService(subjectAdminRepository usecase.SubjectAdminRepository) usecase.SubjectAdminService {
	return &subjectAdminServiceImplementation{
		SubjectAdminRepository: subjectAdminRepository,
	}
}
