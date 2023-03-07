package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type MentorLevelRepository interface {
	Fetch(param string) (data []model.APIMentorLevelResponse, err error)
	FetchByID(id string) (data *model.APIMentorLevelResponse, err error)
	FetchBySlug(slug string) (data *model.APIMentorLevelResponse, err error)
	Create(entity domain.MentorLevel) (e *domain.MentorLevel, err error)
}

type MentorLevelService interface {
	Fetch(param string) (data []model.APIMentorLevelResponse, err error)
	FetchByID(id string) (data *model.APIMentorLevelResponse, err error)
	FetchBySlug(slug string) (data *model.APIMentorLevelResponse, err error)
	Create(request model.CreateMentorLevelRequest) (e *domain.MentorLevel, err error)
}
