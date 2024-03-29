package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
)

type implementSubjectService struct {
	SubjectRepository usecaseMentor.SubjectRepository
}

// GetDataByID implements mentor.SubjectService
func (service *implementSubjectService) GetDataByID(id string) (data *domain.Subject, err error) {
	return service.SubjectRepository.GetDataByID(id)
}

// GetData implements mentor.SubjectService
func (service *implementSubjectService) GetData(param string) (data []domain.Subject, err error) {
	return service.SubjectRepository.GetData(param)
}

func NewSubjectService(subjectRepository usecaseMentor.SubjectRepository) usecaseMentor.SubjectService {
	return &implementSubjectService{
		SubjectRepository: subjectRepository,
	}
}
