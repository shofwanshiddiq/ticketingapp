package services

import (
	"ticketingapp/entity"
	"ticketingapp/repositories"
)

type EventsService interface {
	Create(event *entity.Event) error
	FindByID(id uint) (*entity.Event, error)
	FindAll(page, limit int) ([]entity.Event, int64, error)
	Update(event *entity.Event) error
	Delete(id uint) error
}

type eventsService struct {
	repo repositories.EventsRepository
}

func NewEventsService(repo repositories.EventsRepository) EventsService {
	return &eventsService{repo}
}

func (s *eventsService) Create(event *entity.Event) error {
	return s.repo.Create(event)
}

func (s *eventsService) FindByID(id uint) (*entity.Event, error) {
	return s.repo.FindByID(id)
}

func (s *eventsService) FindAll(page, limit int) ([]entity.Event, int64, error) {
	offset := (page - 1) * limit
	return s.repo.FindAll(offset, limit)
}

func (s *eventsService) Update(event *entity.Event) error {
	return s.repo.Update(event)
}

func (s *eventsService) Delete(id uint) error {
	return s.repo.Delete(id)
}
