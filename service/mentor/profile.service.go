package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	mentorUsecase "github.com/bagusyanuar/go_tb/usecase/mentor"
)

type profileServiceImplementation struct {
	ProfileRepository mentorUsecase.ProfileRepository
}

// GetMyslug implements mentor.ProfileService
func (service *profileServiceImplementation) GetMyslug(id string) (slug string, err error) {
	return service.ProfileRepository.GetMyslug(id)
}

// GetProfile implements mentor.ProfileService
func (service *profileServiceImplementation) GetProfile(id string) (data *domain.User, err error) {
	return service.ProfileRepository.GetProfile(id)
}

func NewProfileService(profileRepository mentorUsecase.ProfileRepository) mentorUsecase.ProfileService {
	return &profileServiceImplementation{
		ProfileRepository: profileRepository,
	}
}
