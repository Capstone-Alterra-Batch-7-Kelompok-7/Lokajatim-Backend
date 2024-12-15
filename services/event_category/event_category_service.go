package event_category

import (
	"lokajatim/entities"
	"lokajatim/repositories/event_category"
)

type EventCategoryServiceImpl struct {
	repo event_category.EventCategoryRepository
}

func NewEventCategoryService(repo event_category.EventCategoryRepository) *EventCategoryServiceImpl {
	return &EventCategoryServiceImpl{repo: repo}
}

func (s *EventCategoryServiceImpl) GetAll() ([]entities.EventCategory, error) {
	return s.repo.GetAll()
}

func (s *EventCategoryServiceImpl) GetByID(id int) (entities.EventCategory, error) {
	return s.repo.GetByID(id)
}

func (s *EventCategoryServiceImpl) Create(category entities.EventCategory) (entities.EventCategory, error) {
	return s.repo.Create(category)
}

func (s *EventCategoryServiceImpl) Update(category entities.EventCategory) (entities.EventCategory, error) {
	existingCategory, err := s.repo.GetByID(category.ID)
	if err != nil {
		return entities.EventCategory{}, err
	}

	existingCategory.Name = category.Name
	return s.repo.Update(existingCategory)
}

func (s *EventCategoryServiceImpl) Delete(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
