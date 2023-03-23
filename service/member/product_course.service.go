package member

import (
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementProductCourseService struct {
	ProductCourseRepository usecaseMember.ProductCourseRepository
}

// Check implements member.ProductCourseService
func (service *implementProductCourseService) Check() (data []model.APIProductCourseResponse, err error) {
	// dt, err := service.ProductCourseRepository.Check()
	// if err != nil {
	// 	return data, err
	// }
	// var d []model.APIProductCourseResponse
	// for _, v := range dt {
	// 	var areas []model.WithDistrictScheme
	// 	for _, vArea := range v.User.Areas {
	// 		areas = append(areas, model.WithDistrictScheme{
	// 			ID:   vArea.ID,
	// 			Name: vArea.Name,
	// 			Code: vArea.Code,
	// 		})
	// 	}
	// 	d = append(d, model.APIProductCourseResponse{
	// 		ProductCourse: model.ProductCourse{
	// 			ID: v.ID,
	// 		},

	// 		User: &model.WithUserScheme{
	// 			ID:   v.User.ID,
	// 			Area: areas,
	// 		},
	// 	})
	// }
	// return d, nil
	return
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
				ID:    productCourse.ID,
				Slug:  productCourse.Slug,
				Price: productCourse.Price,
			},
			Grade: productCourse.Grade.Name,

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

func NewProductCourseService(productCourseRepository usecaseMember.ProductCourseRepository) usecaseMember.ProductCourseService {
	return &implementProductCourseService{
		ProductCourseRepository: productCourseRepository,
	}
}
