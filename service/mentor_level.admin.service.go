package service

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
)

type mentorLevelAdminServiceImplementation struct {
	MentorLevelAdminRepository usecase.MentorLevelAdminRepository
}

// Create implements usecase.MentorLevelAdminService
func (service *mentorLevelAdminServiceImplementation) Create(request domain.CreateMentorLevelRequest) (data *domain.MentorLevel, err error) {
	entity := domain.MentorLevel{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.MentorLevelAdminRepository.Create(entity)
}

// Delete implements usecase.MentorLevelAdminService
func (service *mentorLevelAdminServiceImplementation) Delete(id string) (err error) {
	return service.MentorLevelAdminRepository.Delete(id)
}

// FindAll implements usecase.MentorLevelAdminService
func (service *mentorLevelAdminServiceImplementation) FindAll(param string) (data []domain.MentorLevel, err error) {
	return service.MentorLevelAdminRepository.FindAll(param)
}

// FindByID implements usecase.MentorLevelAdminService
func (service *mentorLevelAdminServiceImplementation) FindByID(id string) (data *domain.MentorLevel, err error) {
	return service.MentorLevelAdminRepository.FindByID(id)
}

// Patch implements usecase.MentorLevelAdminService
func (service *mentorLevelAdminServiceImplementation) Patch(id string, request domain.CreateMentorLevelRequest) (data *domain.MentorLevel, err error) {
	entity := domain.MentorLevel{
		Name: request.Name,
		Slug: common.MakeSlug(request.Name),
	}
	return service.MentorLevelAdminRepository.Patch(id, entity)
}

func NewMentorLevelAdminService(mentorLevelAdminRepository usecase.MentorLevelAdminRepository) usecase.MentorLevelAdminService {
	return &mentorLevelAdminServiceImplementation{
		MentorLevelAdminRepository: mentorLevelAdminRepository,
	}
}
