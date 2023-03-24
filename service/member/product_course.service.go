package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementProductCourseService struct {
	ProductCourseRepository usecaseMember.ProductCourseRepository
}

// GetData implements member.ProductCourseService
func (service *implementProductCourseService) GetData(subjectID string, gradeID string, cityID string) (data []model.APIProductCourseResponse, err error) {
	productCourses, err := service.ProductCourseRepository.GetData(subjectID, gradeID, cityID)
	if err != nil {
		return data, err
	}
	if len(productCourses) <= 0 {
		return []model.APIProductCourseResponse{}, nil
	}
	for _, productCourse := range productCourses {
		data = append(data, model.APIProductCourseResponse{
			ProductCourse: model.ProductCourse{
				ID:          productCourse.ID,
				UserID:      productCourse.UserID,
				SubjectID:   productCourse.SubjectID,
				GradeID:     productCourse.GradeID,
				Slug:        productCourse.Slug,
				Price:       productCourse.Price,
				Method:      productCourse.Method,
				IsAvailable: productCourse.IsAvailable,
			},
			Subject: service.transformSubject(productCourse.Subject),
			Grade:   service.transformGrade(productCourse.Grade),
			User: model.ProductCourseWithUserScheme{
				ID:       productCourse.User.ID,
				Email:    productCourse.User.Email,
				Username: productCourse.User.Username,
				Name:     productCourse.User.Mentor.Name,
			},
		})
	}
	return data, nil
}

func (service *implementProductCourseService) transformGrade(grade *domain.Grade) *model.ProductCourseWithGradeScheme {
	if grade == nil {
		return nil
	}
	return &model.ProductCourseWithGradeScheme{
		ID:   grade.ID,
		Name: grade.Name,
		Slug: grade.Slug,
	}
}

func (service *implementProductCourseService) transformSubject(subject *domain.Subject) *model.ProductCourseWithSubjectScheme {
	if subject == nil {
		return nil
	}
	return &model.ProductCourseWithSubjectScheme{
		ID:   subject.ID,
		Name: subject.Name,
		Slug: subject.Slug,
		Icon: subject.Icon,
	}
}
func NewProductCourseService(productCourseRepository usecaseMember.ProductCourseRepository) usecaseMember.ProductCourseService {
	return &implementProductCourseService{
		ProductCourseRepository: productCourseRepository,
	}
}
