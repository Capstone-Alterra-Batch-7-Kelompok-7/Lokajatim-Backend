package event_category

import "lokajatim/entities"

type EventCategoryRepository interface {
	GetAll() ([]entities.EventCategory, error)
	GetByID(id int) (entities.EventCategory, error)
	Create(category entities.EventCategory) (entities.EventCategory, error)
	Update(category entities.EventCategory) (entities.EventCategory, error)
	Delete(id int) error
}

