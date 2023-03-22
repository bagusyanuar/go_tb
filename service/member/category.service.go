package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
)

type implementCategoryService struct {
	CategoryRepository usecaseMember.CategoryRepository
}

// GetData implements member.CategoryService
func (service *implementCategoryService) GetData(q string) (data []model.APICategoryResponse, err error) {
	categories, err := service.CategoryRepository.GetData(q)
	if err != nil {
		return data, err
	}

	//map to response
	var response []model.APICategoryResponse

	for _, category := range categories {
		tmpCategory := model.APICategoryResponse{
			Category: model.Category{
				ID:   category.ID,
				Name: category.Name,
				Slug: category.Slug,
			},
			Subjects: service.transformSubject(category.Subjects),
		}
		response = append(response, tmpCategory)
	}

	return response, nil
}

func (service *implementCategoryService) transformSubject(subjects []domain.Subject) []model.CategoryWithSubjectScheme {
	var data []model.CategoryWithSubjectScheme
	for _, subject := range subjects {
		data = append(data, model.CategoryWithSubjectScheme{
			ID:   subject.ID,
			Name: subject.Name,
			Slug: subject.Slug,
			Icon: subject.Icon,
		})
	}
	if len(data) > 0 {
		return data
	}
	return []model.CategoryWithSubjectScheme{}
}

func NewCategoryService(categoryRepository usecaseMember.CategoryRepository) usecaseMember.CategoryService {
	return &implementCategoryService{
		CategoryRepository: categoryRepository,
	}
}
