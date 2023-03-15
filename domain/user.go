package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  *string        `json:"password"`
	Roles     datatypes.JSON `json:"roles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Mentor    *Mentor        `gorm:"foreignKey:UserID" json:"mentor"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

func (User) TableName() string {
	return "users"
}

type CreateSignUpMentorRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Avatar   string `json:"avatar"`
}

type CreateSignInMentorRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
