package event

import (
	"lokajatim/entities"
	"lokajatim/repositories/event"
)

type eventService struct {
	repo event.EventRepository
}

func NewEventService(repo event.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) GetAll() ([]entities.Event, error) {
	return s.repo.GetAll()
}

func (s *eventService) GetByID(id uint) (*entities.Event, error) {
	return s.repo.GetByID(id)
}

func (s *eventService) Create(event *entities.Event) error {
	return s.repo.Create(event)
}

func (s *eventService) Update(event *entities.Event) error {
	return s.repo.Update(event)
}

func (s *eventService) Delete(id uint) error {
	return s.repo.Delete(id)
}