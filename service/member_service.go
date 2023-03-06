package service

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type memberServiceImplementation struct {
	MemberRepository usecase.MemberRepository
}

// Create implements usecase.MemberService
func (service *memberServiceImplementation) Create(request model.CreateMemberRequest) (e *domain.Member, err error) {
	entity := domain.Member{
		UserID:  request.UserID,
		Name:    request.Name,
		Phone:   request.Phone,
		Address: request.Address,
	}
	return service.MemberRepository.Create(entity)
}

// Fetch implements usecase.MemberService
func (service *memberServiceImplementation) Fetch() (data []model.APIMemberResponse, err error) {
	return service.MemberRepository.Fetch()
}

// FetchByID implements usecase.MemberService
func (*memberServiceImplementation) FetchByID(id string) (data *model.APIMemberResponse, err error) {
	panic("unimplemented")
}

func NewMemberService(memberRepository usecase.MemberRepository) usecase.MemberService {
	return &memberServiceImplementation{
		MemberRepository: memberRepository,
	}
}
