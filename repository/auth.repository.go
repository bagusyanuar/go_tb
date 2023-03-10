package repository

import (
	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"gorm.io/gorm"
)

type authRepositoryImplementation struct {
	Database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) usecase.AuthRepository {
	return &authRepositoryImplementation{
		Database: database,
	}
}

// SignUpMember implements usecase.AuthRepository
func (repository *authRepositoryImplementation) SignUpMember(user domain.User, member domain.Member) (data *common.JWTSignReturn, err error) {
	//open database transaction
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

	vMember := domain.Member{
		UserID:  user.ID,
		Name:    member.Name,
		Phone:   member.Phone,
		Address: member.Address,
	}
	if err = tx.Create(&vMember).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	res := common.JWTSignReturn{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     "member",
	}
	return &res, nil
}

// SignUpMentor implements usecase.AuthRepository
func (repository *authRepositoryImplementation) SignUpMentor(user domain.User, mentor domain.Mentor) (data *common.JWTSignReturn, err error) {
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

	res := common.JWTSignReturn{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     "mentor",
	}
	return &res, nil
}
