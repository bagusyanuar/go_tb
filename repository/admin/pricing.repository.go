package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementPricingRepository struct {
	Database *gorm.DB
}

// Create implements admin.PricingRepository
func (repository *implementPricingRepository) Create(entity domain.Pricing) (data *domain.Pricing, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewPricingRepository(database *gorm.DB) usecaseAdmin.PricingRepository {
	return &implementPricingRepository{
		Database: database,
	}
}
