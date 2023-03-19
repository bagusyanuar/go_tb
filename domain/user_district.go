package domain

import "github.com/google/uuid"

type UserDistrict struct {
	UserID     uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"user_id"`
	DistrictID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"district_id"`
}

func (UserDistrict) TableName() string {
	return "user_districts"
}
