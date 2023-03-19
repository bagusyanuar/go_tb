package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementAreaRepository struct {
	Database *gorm.DB
}

// AppendAreas implements mentor.AreaRepository
func (repository *implementAreaRepository) AppendAreas(id string, entity []domain.UserDistrict) (err error) {
	tx := repository.Database.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Where("user_id = ?", id).Delete(&domain.UserDistrict{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewAreaRepository(database *gorm.DB) usecaseMentor.AreaRepository {
	return &implementAreaRepository{
		Database: database,
	}
}
