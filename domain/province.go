package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Code      string         `json:"code"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Cities    []City         `gorm:"foreignKey:ProvinceID" json:"cities"`
}

func (province *Province) BeforeCreate(tx *gorm.DB) (err error) {
	province.ID = uuid.New()
	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()
	return
}

func (Province) TableName() string {
	return "provinces"
}

//request to create new category
type CreateProvinceRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
