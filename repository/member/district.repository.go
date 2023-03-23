package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementDistrictRepository struct {
	Database *gorm.DB
}

// GetData implements member.DistrictRepository
func (repository *implementDistrictRepository) GetData(q string) (data []domain.District, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+q+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewDistrictRepository(database *gorm.DB) usecaseMember.DistrictRepository {
	return &implementDistrictRepository{
		Database: database,
	}
}
