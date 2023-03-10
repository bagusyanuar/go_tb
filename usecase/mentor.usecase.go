package usecase

import (
	"github.com/bagusyanuar/go_tb/model"
)

type MentorRepository interface {
	// Fetch() (data []model.APIMemberResponse, err error)
	FetchByID(id string) (data *model.Mentor, err error)
	FetchByUserID(userID string) (data *model.Mentor, err error)
	// Create(entity domain.Member) (e *domain.Member, err error)
}

type MentorService interface {
	// Fetch() (data []model.APIMemberResponse, err error)
	FetchByID(id string) (data *model.Mentor, err error)
	// Create(request model.CreateMemberRequest) (e *domain.Member, err error)
}
