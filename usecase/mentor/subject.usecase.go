package mentor

import "github.com/bagusyanuar/go_tb/domain"

type SubjectRepository interface {
	GetData(param string) (data []domain.Subject, err error)
}

type SubjectService interface {
	GetData(param string) (data []domain.Subject, err error)
}
