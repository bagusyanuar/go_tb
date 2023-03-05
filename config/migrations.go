package config

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email     string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username  string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Member struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"user_id"`
	Name      string         `gorm:"type:varchar(255);" json:"name"`
	Phone     string         `gorm:"type:varchar(25);unique;" json:"phone"`
	Address   string         `gorm:"type:text;" json:"address"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	User      User           `gorm:"foreignKey:UserID"`
}
