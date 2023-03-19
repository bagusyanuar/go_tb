package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/google/uuid"
)

type implementPricingService struct {
	PricingRepository usecaseAdmin.PricingRepository
}

// Create implements admin.PricingService
func (service *implementPricingService) Create(request request.CreatePricingRequest) (data *domain.Pricing, err error) {
	gradeUUID, err := uuid.Parse(request.GradeID)
	if err != nil {
		return nil, err
	}
	cityUUID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}
	mentorLevelUUID, err := uuid.Parse(request.MentorLevelID)
	if err != nil {
		return nil, err
	}

	entity := domain.Pricing{
		GradeID:       gradeUUID,
		CityID:        cityUUID,
		MentorLevelID: mentorLevelUUID,
		Method:        request.Method,
		Price:         request.Price,
	}
	return service.PricingRepository.Create(entity)
}

func NewPricingservice(pricingRepository usecaseAdmin.PricingRepository) usecaseAdmin.PricingService {
	return &implementPricingService{
		PricingRepository: pricingRepository,
	}
}
