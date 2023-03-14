package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type authMentorRepositoryImplementation struct {
	Database *gorm.DB
}

// SignIn implements usecase.AuthMentorRepository
func (*authMentorRepositoryImplementation) SignIn(request model.CreateMentorSignInRequest) (data *domain.User, err error) {
	panic("unimplemented")
}

// SignUp implements usecase.AuthMentorRepository
func (repository *authMentorRepositoryImplementation) SignUp(user domain.User, mentor domain.Mentor) (data *domain.User, err error) {
	tx := repository.Database.Begin()
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

func NewAuthMentorRepository(database *gorm.DB) usecase.AuthMentorRepository {
	return &authMentorRepositoryImplementation{
		Database: database,
	}
}
