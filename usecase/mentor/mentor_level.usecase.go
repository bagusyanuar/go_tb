package mentor

import "github.com/bagusyanuar/go_tb/domain"

type MentorLevelRepository interface {
	GetData(param string) (data []domain.MentorLevel, err error)
	GetDataByID(id string) (data *domain.MentorLevel, err error)
	GetDataBySlug(slug string) (data *domain.MentorLevel, err error)
}

type MentorLevelService interface {
	GetData(param string) (data []domain.MentorLevel, err error)
	GetDataByID(id string) (data *domain.MentorLevel, err error)
	GetDataBySlug(slug string) (data *domain.MentorLevel, err error)
}
