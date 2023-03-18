package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"gorm.io/gorm"
)

type implementAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements member.AuthRepository
func (repository *implementAuthRepository) SignIn(user domain.User) (data *domain.User, err error) {
	var entity domain.User
	if err = repository.Database.Debug().
		Preload("Mentor").
		Joins("JOIN members ON users.id = members.user_id").
		Where("email = ?", user.Email).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// SignUp implements member.AuthRepository
func (repository *implementAuthRepository) SignUp(user domain.User, member domain.Member) (data *domain.User, err error) {
	tx := repository.Database.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	entityMember := domain.Member{
		UserID:  user.ID,
		Name:    member.Name,
		Phone:   member.Phone,
		Address: member.Address,
	}
	if err = tx.Create(&entityMember).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &user, nil
}

func NewAuthRepository(database *gorm.DB) usecaseMember.AuthRepository {
	return &implementAuthRepository{
		Database: database,
	}
}
