package category

import (
	"sagala/factory"
	"sagala/internal/dto"
	"sagala/internal/models"
	"sagala/internal/repository"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(factory *factory.Factory) *CategoryService {
	return &CategoryService{
		categoryRepo: factory.CategoryRepo,
	}
}

func (s *CategoryService) Create(request *dto.CategoryRequest) error {
	category := models.Category{
		Name:      request.Name,
		SortOrder: request.SortOrder,
	}

	if err := s.categoryRepo.Create(&category); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) Update(id uint, request *dto.CategoryRequest) error {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	category.Name = request.Name
	category.SortOrder = request.SortOrder

	return s.categoryRepo.Update(&category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}

func (s *CategoryService) Index() ([]dto.CategoryResponse, error) {
	var resp dto.CategoryResponse
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return resp.BuildListCategoryResponse(categories), nil
}

func (s *CategoryService) Show(id uint) (*dto.CategoryResponse, error) {
	var resp dto.CategoryResponse
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return resp.BuildCategoryResponse(category), nil
}
