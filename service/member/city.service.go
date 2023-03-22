package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementCityService struct {
	CityRepository usecaseMember.CityRepository
}

// GetData implements member.CityService
func (service *implementCityService) GetData(q string) (data []model.APICityResponse, err error) {
	cities, err := service.CityRepository.GetData(q)
	if err != nil {
		return data, err
	}

	for _, city := range cities {
		apiCity := model.APICityResponse{
			City: model.City{
				ID:   city.ID,
				Name: city.Name,
				Code: city.Code,
			},
			Province:  service.transformProvince(city.Province),
			Districts: service.transformDistricts(city.Districts),
		}
		data = append(data, apiCity)
	}
	return data, nil
}

func (service *implementCityService) transformProvince(province *domain.Province) *model.CityWithProvinceScheme {
	if province == nil {
		return nil
	}
	return &model.CityWithProvinceScheme{
		ID:   province.ID,
		Name: province.Name,
		Code: province.Code,
	}
}

func (service *implementCityService) transformDistricts(districts []domain.District) []model.CityWithDistrictScheme {
	if len(districts) <= 0 {
		return []model.CityWithDistrictScheme{}
	}
	var data []model.CityWithDistrictScheme
	for _, district := range districts {
		data = append(data, model.CityWithDistrictScheme{
			ID:   district.ID,
			Name: district.Name,
			Code: district.Code,
		})
	}
	return data
}

func NewCityService(cityRepository usecaseMember.CityRepository) usecaseMember.CityService {
	return &implementCityService{
		CityRepository: cityRepository,
	}
}
