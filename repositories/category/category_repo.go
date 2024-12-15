package category

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepositoryImpl{db: db}
}

func (r *CategoryRepositoryImpl) GetCategories() ([]entities.Category, error) {
	var categories []entities.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *CategoryRepositoryImpl) GetCategoryByID(id int) (entities.Category, error) {
	var category entities.Category
	result := r.db.First(&category, id)
	if result.Error != nil {
		return entities.Category{}, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) CreateCategory(category entities.Category) (entities.Category, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return entities.Category{}, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(id int, category entities.Category) (entities.Category, error) {
	if err := r.db.Model(&entities.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return entities.Category{}, err
	}

	var updatedCategory entities.Category
	if err := r.db.First(&updatedCategory, id).Error; err != nil {
		return entities.Category{}, err
	}
	return updatedCategory, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(id int) error{
	var category entities.Category
	if err := r.db.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}
