package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"gorm.io/gorm"
)

type implementAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements mentor.AuthMentorRepository
func (repository *implementAuthRepository) SignIn(user domain.User) (data *domain.User, err error) {
	var entity domain.User
	if err = repository.Database.Debug().
		Preload("Mentor").
		Joins("JOIN mentors ON users.id = mentors.user_id").
		Where("email = ?", user.Email).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// SignUp implements mentor.AuthMentorRepository
func (repository *implementAuthRepository) SignUp(user domain.User, mentor domain.Mentor) (data *domain.User, err error) {
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

	entityMentor := domain.Mentor{
		UserID:        user.ID,
		MentorLevelID: mentor.MentorLevelID,
		Name:          mentor.Name,
		Slug:          mentor.Slug,
		Gender:        mentor.Gender,
		Phone:         mentor.Phone,
		Address:       mentor.Address,
		Avatar:        mentor.Avatar,
	}
	if err = tx.Create(&entityMentor).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &user, nil
}

func NewAuthRepository(database *gorm.DB) usecaseMentor.AuthRepository {
	return &implementAuthRepository{
		Database: database,
	}
}
