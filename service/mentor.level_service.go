package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type mentorLevelServiceImplementation struct {
	MentorLevelRepository usecase.MentorLevelRepository
}

func NewMentorLevelService(mentorLevelRepository usecase.MentorLevelRepository) usecase.MentorLevelService {
	return &mentorLevelServiceImplementation{
		MentorLevelRepository: mentorLevelRepository,
	}
}

// Create implements usecase.MentorLevelService
func (service *mentorLevelServiceImplementation) Create(request model.CreateMentorLevelRequest) (e *domain.MentorLevel, err error) {
	entity := domain.MentorLevel{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.MentorLevelRepository.Create(entity)
}

// Fetch implements usecase.MentorLevelService
func (service *mentorLevelServiceImplementation) Fetch(param string) (data []model.APIMentorLevelResponse, err error) {
	return service.MentorLevelRepository.Fetch(param)
}

// FetchByID implements usecase.MentorLevelService
func (service *mentorLevelServiceImplementation) FetchByID(id string) (data *model.APIMentorLevelResponse, err error) {
	return service.MentorLevelRepository.FetchByID(id)
}

// FetchBySlug implements usecase.MentorLevelService
func (service *mentorLevelServiceImplementation) FetchBySlug(slug string) (data *model.APIMentorLevelResponse, err error) {
	return service.MentorLevelRepository.FetchBySlug(slug)
}
