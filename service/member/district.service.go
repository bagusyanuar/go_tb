package member

import (
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementDistrictService struct {
	DistrictRepository usecaseMember.DistrictRepository
}

// GetData implements member.DistrictService
func (service *implementDistrictService) GetData(q string) (data []model.District, err error) {
	districts, err := service.DistrictRepository.GetData(q)
	if err != nil {
		return data, nil
	}
	for _, district := range districts {
		data = append(data, model.District{
			ID:     district.ID,
			CityID: district.CityID,
			Name:   district.Name,
			Code:   district.Code,
		})
	}
	return data, nil
}

func NewDistrictService(districtRepository usecaseMember.DistrictRepository) usecaseMember.DistrictService {
	return &implementDistrictService{
		DistrictRepository: districtRepository,
	}
}
