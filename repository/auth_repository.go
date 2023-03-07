package repository

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
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
func (repository *authRepositoryImplementation) SignUpMember(user domain.User, member domain.Member) (data *model.APISignUpResponse, err error) {
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

	res := model.APISignUpResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     "member",
	}
	return &res, nil
}
