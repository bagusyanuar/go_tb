package member

import "github.com/bagusyanuar/go_tb/domain"

type SubjectRepository interface {
	GetData(q string) (data []domain.Subject, err error)
}
