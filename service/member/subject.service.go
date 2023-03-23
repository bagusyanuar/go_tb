package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementSubjectService struct {
	SubjectRepository usecaseMember.SubjectRepository
}

// GetDataByID implements member.SubjectService
func (service *implementSubjectService) GetDataByID(id string) (data *model.APISubjectResponse, err error) {
	subject, err := service.SubjectRepository.GetDataByID(id)
	if err != nil {
		return nil, err
	}

	data = &model.APISubjectResponse{
		Subject: model.Subject{
			ID:         subject.ID,
			CategoryID: subject.CategoryID,
			Name:       subject.Name,
			Icon:       subject.Icon,
			Slug:       subject.Slug,
		},
		Grades: service.transformGrades(subject.Grades),
	}
	return data, nil
}

// GetData implements member.SubjectService
func (service *implementSubjectService) GetData(q string) (data []model.Subject, err error) {
	subjects, err := service.SubjectRepository.GetData(q)
	if err != nil {
		return data, err
	}
	for _, subject := range subjects {
		data = append(data, model.Subject{
			ID:         subject.ID,
			CategoryID: subject.CategoryID,
			Name:       subject.Name,
			Slug:       subject.Slug,
			Icon:       subject.Icon,
		})
	}
	return data, nil
}

func (service *implementSubjectService) transformGrades(grades []domain.Grade) []model.SubjectWithGradeScheme {
	if len(grades) <= 0 {
		return []model.SubjectWithGradeScheme{}
	}
	var data []model.SubjectWithGradeScheme
	for _, grade := range grades {
		data = append(data, model.SubjectWithGradeScheme{
			ID:   grade.ID,
			Name: grade.Name,
			Slug: grade.Slug,
		})
	}
	return data
}

func NewSubjectService(subjectRepository usecaseMember.SubjectRepository) usecaseMember.SubjectService {
	return &implementSubjectService{
		SubjectRepository: subjectRepository,
	}
}
