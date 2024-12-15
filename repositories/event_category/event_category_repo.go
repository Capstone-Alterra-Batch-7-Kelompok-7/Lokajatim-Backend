package event_category

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type EventCategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewEventCategoryRepository(db *gorm.DB) EventCategoryRepository {
	return &EventCategoryRepositoryImpl{db: db}
}

func (r *EventCategoryRepositoryImpl) GetAll() ([]entities.EventCategory, error) {
	var categories []entities.EventCategory
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *EventCategoryRepositoryImpl) GetByID(id int) (entities.EventCategory, error) {
	var category entities.EventCategory
	if err := r.db.First(&category, id).Error; err != nil {
		return entities.EventCategory{}, err
	}
	return category, nil
}

func (r *EventCategoryRepositoryImpl) Create(category entities.EventCategory) (entities.EventCategory, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return entities.EventCategory{}, err
	}
	return category, nil
}

func (r *EventCategoryRepositoryImpl) Update(category entities.EventCategory) (entities.EventCategory, error) {
	if err := r.db.Save(&category).Error; err != nil {
		return entities.EventCategory{}, err
	}
	return category, nil
}

func (r *EventCategoryRepositoryImpl) Delete(id int) error {
	if err := r.db.Delete(&entities.EventCategory{}, id).Error; err != nil {
		return err
	}
	return nil
}