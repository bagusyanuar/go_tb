package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"gorm.io/gorm"
)

type implementCategoryRepository struct {
	Database *gorm.DB
}

// Create implements admin.CategoryRepository
func (repository *implementCategoryRepository) Create(entity domain.Category) (data *domain.Category, err error) {
	panic("unimplemented")
}

// Delete implements admin.CategoryRepository
func (repository *implementCategoryRepository) Delete(id string) (err error) {
	panic("unimplemented")
}

// FindAll implements admin.CategoryRepository
func (repository *implementCategoryRepository) FindAll(param string) (data []domain.Category, err error) {
	panic("unimplemented")
}

// FindByID implements admin.CategoryRepository
func (repository *implementCategoryRepository) FindByID(id string) (data *domain.Category, err error) {
	panic("unimplemented")
}

// Patch implements admin.CategoryRepository
func (repository *implementCategoryRepository) Patch(id string, entity domain.Category) (data *domain.Category, err error) {
	panic("unimplemented")
}

func NewCategoryRepository(database *gorm.DB) usecaseAdmin.CategoryRepository {
	return &implementCategoryRepository{
		Database: database,
	}
}
