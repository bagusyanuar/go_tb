package mentor

import "github.com/bagusyanuar/go_tb/domain"

type GradeRepository interface {
	GetData(param string) (data []domain.Grade, err error)
	GetDataByID(id string) (data *domain.Grade, err error)
	GetDataBySlug(slug string) (data *domain.Grade, err error)
}

type GradeService interface {
	GetData(param string) (data []domain.Grade, err error)
	GetDataByID(id string) (data *domain.Grade, err error)
	GetDataBySlug(slug string) (data *domain.Grade, err error)
}
