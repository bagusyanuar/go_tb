package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type AreaRepository interface {
	AppendAreas(id string, entity []domain.UserDistrict) (err error)
}

type AreaService interface {
	AppendAreas(id string, request request.CreateMentorAreaRequest) (err error)
}
