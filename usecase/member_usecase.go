package usecase

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
)

type MemberRepository interface {
	Fetch() (data []model.APIMemberResponse, err error)
	FetchByID(id string) (data *model.APIMemberResponse, err error)
	Create(entity domain.Member) (e *domain.Member, err error)
}

type MemberService interface {
	Fetch() (data []model.APIMemberResponse, err error)
	FetchByID(id string) (data *model.APIMemberResponse, err error)
	Create(request model.CreateMemberRequest) (e *domain.Member, err error)
}
