package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
)

type implementMentorLevelService struct {
	MentorLevelRepository usecaseAdmin.MentorLevelAdminRepository
}

// Create implements admin.MentorLevelAdminService
func (service *implementMentorLevelService) Create(request request.CreateMentorLevelRequest) (data *domain.MentorLevel, err error) {
	panic("unimplemented")
}

// Delete implements admin.MentorLevelAdminService
func (service *implementMentorLevelService) Delete(id string) (err error) {
	panic("unimplemented")
}

// FindAll implements admin.MentorLevelAdminService
func (service *implementMentorLevelService) FindAll(param string) (data []domain.MentorLevel, err error) {
	panic("unimplemented")
}

// FindByID implements admin.MentorLevelAdminService
func (service *implementMentorLevelService) FindByID(id string) (data *domain.MentorLevel, err error) {
	panic("unimplemented")
}

// Patch implements admin.MentorLevelAdminService
func (service *implementMentorLevelService) Patch(id string, request request.CreateMentorLevelRequest) (data *domain.MentorLevel, err error) {
	panic("unimplemented")
}

func NewMentorLevelService(mentorLevelRepository usecaseAdmin.MentorLevelAdminRepository) usecaseAdmin.MentorLevelAdminService {
	return &implementMentorLevelService{
		MentorLevelRepository: mentorLevelRepository,
	}
}
