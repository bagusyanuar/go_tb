package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementProductCourseService struct {
	ProductCourseRepository usecaseMember.ProductCourseRepository
}

// GetData implements member.ProductCourseService
func (service *implementProductCourseService) GetData(subjectID string, gradeID string, cityID string) (data []domain.ProductCourse, err error) {
	return service.ProductCourseRepository.GetData(subjectID, gradeID, cityID)
}

func NewProductCourseService(productCourseRepository usecaseMember.ProductCourseRepository) usecaseMember.ProductCourseService {
	return &implementProductCourseService{
		ProductCourseRepository: productCourseRepository,
	}
}
