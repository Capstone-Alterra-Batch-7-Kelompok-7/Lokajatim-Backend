package event

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetAll() ([]entities.Event, error) {
	var events []entities.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) GetByID(id uint) (*entities.Event, error) {
	var event entities.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *eventRepository) Create(event *entities.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Update(event *entities.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Event{}, id).Error
}
