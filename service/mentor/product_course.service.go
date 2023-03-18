package mentor

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
)

type implementProductCourseService struct {
	ProductCourseRpository usecaseMentor.ProductCourseRepository
}

// CreateMyProductCourse implements mentor.ProductCourseService
func (service *implementProductCourseService) CreateMyProductCourse(userID string, request request.CreateMyProductCourseRequest) (data *domain.ProductCourse, err error) {
	// userUUID, err := uuid.Parse(userID)
	// if err != nil {
	// 	return nil, err
	// }

	// subjectUUID, err := uuid.Parse(request.SubjectID)
	// if err != nil {
	// 	return nil, err
	// }

	// subjectUUID, err := uuid.Parse(request.SubjectID)
	// if err != nil {
	// 	return nil, err
	// }
	return
}

// GetMyProductCourse implements mentor.ProductCourseService
func (*implementProductCourseService) GetMyProductCourse(id string) (data []domain.ProductCourse, err error) {
	panic("unimplemented")
}

func NewProductCourseService(productCourseRpository usecaseMentor.ProductCourseRepository) usecaseMentor.ProductCourseService {
	return &implementProductCourseService{
		ProductCourseRpository: productCourseRpository,
	}
}
