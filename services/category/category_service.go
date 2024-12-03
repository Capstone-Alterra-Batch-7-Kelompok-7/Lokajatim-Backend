package category

import (
	"lokajatim/entities"
	"lokajatim/repositories/category"
)

type CategoryService struct {
	categoryRepository category.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepo category.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepo}
}

func (s *CategoryService) GetCategories() ([]entities.Category, error) {
	return s.categoryRepository.GetCategories()
}

func (s *CategoryService) GetCategoryByID(id int) (entities.Category, error) {
	return s.categoryRepository.GetCategoryByID(id)
}

func (s *CategoryService) CreateCategory(category entities.Category) (entities.Category, error) {
	return s.categoryRepository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(id int, category entities.Category) (entities.Category, error) {
	return s.categoryRepository.UpdateCategory(id, category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.categoryRepository.DeleteCategory(id)
}
