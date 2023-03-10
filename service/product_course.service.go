package service

import (
	"encoding/json"
	"fmt"

	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
)

type productCourseServiceImplementation struct {
	ProductCourseRepository usecase.ProductCourseRepository
	MentorRepository        usecase.MentorRepository
	SubjectRepository       usecase.SubjectRepository
	GradeRepository         usecase.GradeRepository
}

func NewProductCourseService(
	productCourseRepository usecase.ProductCourseRepository,
	mentorRepository usecase.MentorRepository,
	subjectRepository usecase.SubjectRepository,
	gradeRepository usecase.GradeRepository,
) usecase.ProductCourseService {
	return &productCourseServiceImplementation{
		ProductCourseRepository: productCourseRepository,
		MentorRepository:        mentorRepository,
		SubjectRepository:       subjectRepository,
		GradeRepository:         gradeRepository,
	}
}

// Create implements usecase.ProductCourseService
func (service *productCourseServiceImplementation) Create(request model.CreateProductCourseRequest) (e *domain.ProductCourse, err error) {
	mentor, errFetchMentor := service.MentorRepository.FetchByUserID(request.UserID)
	if errFetchMentor != nil {
		return nil, errFetchMentor
	}
	subject, errFetchSubject := service.SubjectRepository.FetchByID(request.SubjectID)
	if errFetchMentor != nil {
		return nil, errFetchSubject
	}
	grade, errFetchGrade := service.GradeRepository.FetchByID(request.GradeID)
	if errFetchMentor != nil {
		return nil, errFetchGrade
	}

	slug := fmt.Sprintf("%s-%s-%s", subject.Slug, grade.Slug, mentor.Slug)
	methods, _ := json.Marshal(request.Method)
	entity := domain.ProductCourse{
		UserID:      mentor.UserID,
		SubjectID:   subject.ID,
		GradeID:     grade.ID,
		Method:      methods,
		Slug:        slug,
		IsAvailable: true,
	}
	return service.ProductCourseRepository.Create(entity)
}

// Fetch implements usecase.ProductCourseService
func (service *productCourseServiceImplementation) Fetch(param string) (data []model.APIProductCourseResponse, err error) {
	return service.ProductCourseRepository.Fetch(param)
}
