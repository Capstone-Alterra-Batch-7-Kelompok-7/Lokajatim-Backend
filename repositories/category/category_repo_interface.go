package category

import (
	"lokajatim/entities"
)

type CategoryRepositoryInterface interface {
	GetCategories() ([]entities.Category, error)
	GetCategoryByID(id int) (entities.Category, error)
	CreateCategory(category entities.Category) (entities.Category, error)
	UpdateCategory(id int, category entities.Category) (entities.Category, error)
	DeleteCategory(id int) error
}
