package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementCategoryRepository struct {
	Database *gorm.DB
}

// GetData implements member.CategoryRepository
func (repository *implementCategoryRepository) GetData(q string) (data []domain.Category, err error) {
	if err = repository.Database.Debug().Preload("Subjects").Where("name LIKE ?", "%"+q+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewCategoryRepository(database *gorm.DB) usecaseMember.CategoryRepository {
	return &implementCategoryRepository{
		Database: database,
	}
}
