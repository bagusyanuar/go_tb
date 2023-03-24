package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRequest struct {
	ID          uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	StudentID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_student_id;not null;" json:"student_id"`
	MentorID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_mentor_id;not null;" json:"mentor_id"`
	SubjectID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_subject_id;not null;" json:"subject_id"`
	GradeID     uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_grade_id;not null;" json:"grade_id"`
	DistrictID  *uuid.UUID     `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_district_id;" json:"district_id"`
	ReferenceID string         `gorm:"type:varchar(255);not null;unique" json:"reference_id"`
	Method      uint8          `gorm:"not null;" json:"method"`
	Duration    uint8          `gorm:"not null;default:0;" json:"duration"`
	Attendees   uint8          `gorm:"not null;default:0;" json:"attendees"`
	Encounter   uint8          `gorm:"not null;default:0;" json:"encounter"`
	FirstMeet   time.Time      `gorm:"not null;" json:"first_meet"`
	Address     string         `gorm:"type:text;" json:"address"`
	Note        string         `gorm:"type:text;" json:"note"`
	SubTotal    uint           `gorm:"type:int(11);not null;default:0;" json:"sub_total"`
	Discount    uint           `gorm:"type:int(11);not null;default:0;" json:"discount"`
	Total       uint           `gorm:"type:int(11);not null;default:0;" json:"total"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Student     User           `gorm:"foreignKey:StudentID"`
	Mentor      User           `gorm:"foreignKey:MentorID"`
	Subject     Subject        `gorm:"foreignKey:SubjectID"`
	Grade       Grade          `gorm:"foreignKey:GradeID"`
	District    *District      `gorm:"foreignKey:DistrictID"`
}

func (courseRequest *CourseRequest) BeforeCreate(tx *gorm.DB) (err error) {
	courseRequest.ID = uuid.New()
	courseRequest.CreatedAt = time.Now()
	courseRequest.UpdatedAt = time.Now()
	return
}

func (CourseRequest) TableName() string {
	return "course_requests"
}
