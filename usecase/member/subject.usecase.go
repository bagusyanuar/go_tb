package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type SubjectRepository interface {
	GetData(q string) (data []domain.Subject, err error)
	GetDataByID(id string) (data *domain.Subject, err error)
}

type SubjectService interface {
	GetData(q string) (data []model.Subject, err error)
	GetDataByID(id string) (data *model.APISubjectResponse, err error)
}
