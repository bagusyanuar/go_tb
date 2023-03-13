package usecase

import "github.com/bagusyanuar/go_tb/domain"

type MentorLevelAdminRepository interface {
	FindAll(param string) (data []domain.MentorLevel, err error)
	FindByID(id string) (data *domain.MentorLevel, err error)
	Create(entity domain.MentorLevel) (data *domain.MentorLevel, err error)
	Patch(id string, entity domain.MentorLevel) (data *domain.MentorLevel, err error)
	Delete(id string) (err error)
}

type MentorLevelAdminService interface {
	FindAll(param string) (data []domain.MentorLevel, err error)
	FindByID(id string) (data *domain.MentorLevel, err error)
	Create(request domain.CreateMentorLevelRequest) (data *domain.MentorLevel, err error)
	Patch(id string, request domain.CreateMentorLevelRequest) (data *domain.MentorLevel, err error)
	Delete(id string) (err error)
}
