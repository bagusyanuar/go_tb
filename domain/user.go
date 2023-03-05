package domain

import (
	"time"

	"github.com/bagusyanuar/go_tb/model"
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

type UserService interface {
	Fetch() (data []model.CreateUserResponse, err error)
	FetchByID(id string) (*User, error)
	Create(request model.CreateUserRequest) (user *User, err error)
}

type UserRepository interface {
	Fetch() (data []model.CreateUserResponse, err error)
	FetchByID(id string) (result *User, err error)
	Create(user User) (result *User, err error)
}
