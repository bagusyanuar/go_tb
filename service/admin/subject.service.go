package admin

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/google/uuid"
)

type implementSubjectService struct {
	SubjectRepository usecaseAdmin.SubjectRepository
}

// AppendGrade implements admin.SubjectService
func (service *implementSubjectService) AppendGrade(id string, request request.CreateSubjectAppendGradeRequest) (err error) {
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
	return service.SubjectRepository.AppendGrade(id, entity)
}

// Create implements admin.SubjectService
func (service *implementSubjectService) Create(request request.CreateSubjectRequest) (data *domain.Subject, err error) {
	categoryID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return nil, err
	}
	entity := domain.Subject{
		Name:       request.Name,
		Slug:       common.MakeSlug(request.Name),
		CategoryID: categoryID,
	}
	return service.SubjectRepository.Create(entity)
}

// Delete implements admin.SubjectService
func (service *implementSubjectService) Delete(id string) (err error) {
	return service.SubjectRepository.Delete(id)
}

// FindAll implements admin.SubjectService
func (service *implementSubjectService) FindAll(param string) (data []domain.Subject, err error) {
	return service.SubjectRepository.FindAll(param)
}

// FindByID implements admin.SubjectService
func (service *implementSubjectService) FindByID(id string) (data *domain.Subject, err error) {
	return service.SubjectRepository.FindByID(id)
}

// Patch implements admin.SubjectService
func (service *implementSubjectService) Patch(id string, request request.CreateSubjectRequest) (data *domain.Subject, err error) {
	categoryID, err := uuid.Parse(request.CategoryID)
	if err != nil {
		return nil, err
	}
	entity := domain.Subject{
		Name:       request.Name,
		Slug:       common.MakeSlug(request.Name),
		CategoryID: categoryID,
	}
	return service.SubjectRepository.Patch(id, entity)
}

func NewSubjectService(subjectRepository usecaseAdmin.SubjectRepository) usecaseAdmin.SubjectService {
	return &implementSubjectService{
		SubjectRepository: subjectRepository,
	}
}
