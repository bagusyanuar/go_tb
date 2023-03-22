package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementSubjectRepository struct {
	Database *gorm.DB
}

// GetData implements member.SubjectRepository
func (repository *implementSubjectRepository) GetData(q string) (data []domain.Subject, err error) {
	if err = repository.Database.Debug().Where("name LIKE ?", "%"+q+"%").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewSubjectRepository(database *gorm.DB) usecaseMember.SubjectRepository {
	return &implementSubjectRepository{
		Database: database,
	}
}
