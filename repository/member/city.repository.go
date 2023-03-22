package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementCityRepository struct {
	Database *gorm.DB
}

// GetData implements member.CityRepository
func (repository *implementCityRepository) GetData(q string) (data []domain.City, err error) {
	if err = repository.Database.Debug().Preload("Province").Preload("Districts").Where("name LIKE ?", "%"+q+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewCityRepository(database *gorm.DB) usecaseMember.CityRepository {
	return &implementCityRepository{
		Database: database,
	}
}
