package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

func NewMemberRepository(database *gorm.DB) usecase.MemberRepository {
	return &memberRepositoryImplementation{
		Database: database,
	}
}

type memberRepositoryImplementation struct {
	Database *gorm.DB
}

// Create implements usecase.MemberRepository
func (repository *memberRepositoryImplementation) Create(entity domain.Member) (e *domain.Member, err error) {
	if err = repository.Database.Debug().Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Fetch implements usecase.MemberRepository
func (repository *memberRepositoryImplementation) Fetch() (data []model.APIMemberResponse, err error) {
	var entity []domain.Member
	if err = repository.Database.Debug().Model(&entity).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FetchByID implements usecase.MemberRepository
func (*memberRepositoryImplementation) FetchByID(id string) (data *model.APIMemberResponse, err error) {
	panic("unimplemented")
}
