package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"github.com/google/uuid"
)

type implementAreaService struct {
	AreaRepository usecaseMentor.AreaRepository
}

// AppendAreas implements mentor.AreaService
func (service *implementAreaService) AppendAreas(id string, request request.CreateMentorAreaRequest) (err error) {
	var districtIDS []uuid.UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	for _, districtID := range request.DistrictIDS {
		districtUUID, err := uuid.Parse(districtID)
		if err != nil {
			return err
		}
		districtIDS = append(districtIDS, districtUUID)
	}

	entity := []domain.UserDistrict{}
	for _, districtUUID := range districtIDS {
		entity = append(entity, domain.UserDistrict{
			UserID:     userID,
			DistrictID: districtUUID,
		})
	}
	return service.AreaRepository.AppendAreas(id, entity)
}

func NewAreaService(areaRepository usecaseMentor.AreaRepository) usecaseMentor.AreaService {
	return &implementAreaService{
		AreaRepository: areaRepository,
	}
}
