package config

import (
	"database/sql/driver"
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
	Districts []District     `gorm:"many2many:user_districts" json:"districts"`
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

type Category struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Slug      string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Subject struct {
	ID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CategoryID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"category_id"`
	Name       string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Slug       string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	Icon       string         `gorm:"column:icon;type:text" json:"icon"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Category   Category       `gorm:"foreignKey:CategoryID"`
	Grades     []Grade        `gorm:"many2many:subject_grades" json:"grades"`
}

type Grade struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Slug      string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type MentorLevel struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Slug      string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Province struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Code      string         `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type City struct {
	ID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProvinceID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"province_id"`
	Name       string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Code       string         `gorm:"column:code;type:char(6);index:idx_code,unique;" json:"code"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Province   *Province      `gorm:"foreignKey:ProvinceID"`
}

type District struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CityID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"city_id"`
	Name      string         `gorm:"column:name;type:varchar(255);" json:"name"`
	Code      string         `gorm:"column:code;type:char(8);index:idx_code,unique;" json:"code"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	City      *City          `gorm:"foreignKey:CityID"`
}

type genderType string

const (
	MEN   genderType = "men"
	WOMEN genderType = "women"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}

func (gt genderType) Value() (driver.Value, error) {
	return string(gt), nil
}

type Mentor struct {
	ID            uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id,unique;" json:"user_id"`
	MentorLevelID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"mentor_level_id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Slug          string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	Gender        genderType     `gorm:"type:enum('men', 'women')" json:"gender"`
	Phone         string         `gorm:"type:varchar(25);unique;" json:"phone"`
	Address       string         `gorm:"type:text;" json:"address"`
	Avatar        string         `gorm:"type:text;" json:"avatar"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	User          User           `gorm:"foreignKey:UserID"`
	MentorLevel   MentorLevel    `gorm:"foreignKey:MentorLevelID"`
}

type SubjectGrade struct {
	SubjectID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"subject_id"`
	GradeID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"grade_id"`
}

type ProductCourse struct {
	ID          uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID      uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_user_id;" json:"user_id"`
	SubjectID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_subject_id;" json:"subject_id"`
	GradeID     uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_grade_id;" json:"grade_id"`
	Method      datatypes.JSON `gorm:"type:longtext;not null" json:"method"`
	Slug        string         `gorm:"column:slug;index:idx_slug,unique;not null" json:"slug"`
	IsAvailable bool           `gorm:"column:is_available;default:0;not null" json:"is_available"`
	User        User           `gorm:"foreignKey:UserID"`
	Subject     Subject        `gorm:"foreignKey:SubjectID"`
	Grade       Grade          `gorm:"foreignKey:GradeID"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Pricing struct {
	ID            uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	GradeID       uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_grade_id;" json:"grade_id"`
	CityID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;" json:"city_id"`
	MentorLevelID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_mentor_level_id;" json:"mentor_level_id"`
	Method        uint           `gorm:"type:int(11);not null;" json:"method"`
	Price         uint           `gorm:"type:int(11);not null;default:0;" json:"price"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Grade         *Grade         `gorm:"foreignKey:GradeID"`
	City          *City          `gorm:"foreignKey:CityID"`
	MentorLevel   *MentorLevel   `gorm:"foreignKey:MentorLevelID"`
}

type CourseRequest struct {
	ID          uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	StudentID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_student_id;not null;" json:"student_id"`
	MentorID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_mentor_id;not null;" json:"mentor_id"`
	SubjectID   uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_subject_id;not null;" json:"subject_id"`
	GradeID     uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_grade_id;not null;" json:"grade_id"`
	DistrictID  uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_district_id;" json:"district_id"`
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
