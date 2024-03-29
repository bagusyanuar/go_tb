package mentor

import (
	"encoding/json"
	"fmt"

	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"github.com/google/uuid"
)

type implementProductCourseService struct {
	ProductCourseRpository usecaseMentor.ProductCourseRepository
	ProfileRepository      usecaseMentor.ProfileRepository
	SubjectRepository      usecaseMentor.SubjectRepository
	GradeRepository        usecaseMentor.GradeRepository
}

// CreateMyProductCourse implements mentor.ProductCourseService
func (service *implementProductCourseService) CreateMyProductCourse(userID string, request request.CreateMyProductCourseRequest) (data *domain.ProductCourse, err error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	subjectUUID, err := uuid.Parse(request.SubjectID)
	if err != nil {
		return nil, err
	}

	gradeUUID, err := uuid.Parse(request.GradeID)
	if err != nil {
		return nil, err
	}

	slugSubject, err := service.SubjectRepository.GetDataByID(request.SubjectID)
	if err != nil {
		return nil, err
	}

	slugGrade, err := service.GradeRepository.GetDataByID(request.GradeID)
	if err != nil {
		return nil, err
	}

	slugProfile, err := service.ProfileRepository.GetMyslug(userID)
	if err != nil {
		return nil, err
	}

	slug := fmt.Sprintf("%s-%s-%s", slugSubject.Slug, slugGrade.Slug, slugProfile)
	methods, _ := json.Marshal(request.Method)
	entity := domain.ProductCourse{
		UserID:      userUUID,
		SubjectID:   subjectUUID,
		GradeID:     gradeUUID,
		Slug:        slug,
		IsAvailable: true,
		Method:      methods,
	}
	return service.ProductCourseRpository.CreateMyProductCourse(entity)
}

// GetMyProductCourse implements mentor.ProductCourseService
func (service *implementProductCourseService) GetMyProductCourse(id string) (data []domain.ProductCourse, err error) {
	return service.ProductCourseRpository.GetMyProductCourse(id)
}

func NewProductCourseService(productCourseRpository usecaseMentor.ProductCourseRepository, profileRepository usecaseMentor.ProfileRepository, subjectRepository usecaseMentor.SubjectRepository, gradeRepository usecaseMentor.GradeRepository) usecaseMentor.ProductCourseService {
	return &implementProductCourseService{
		ProductCourseRpository: productCourseRpository,
		ProfileRepository:      profileRepository,
		SubjectRepository:      subjectRepository,
		GradeRepository:        gradeRepository,
	}
}
