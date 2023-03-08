package domain

import "github.com/google/uuid"

type SubjectGrade struct {
	SubjectID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"subject_id"`
	GradeID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"grade_id"`
}

func (SubjectGrade) TableName() string {
	return "subject_grades"
}
