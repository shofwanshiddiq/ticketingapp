package repositories

import (
	"ticketingapp/entity"

	"gorm.io/gorm"
)

type EventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) *EventsRepository {
	return &EventsRepository{db: db}
}

func (r *EventsRepository) Create(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *EventsRepository) FindByID(id uint) (*entity.Event, error) {
	var event entity.Event
	err := r.db.Where("id = ?", id).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventsRepository) FindAll(offset, limit int) ([]entity.Event, int64, error) {
	var events []entity.Event
	var total int64

	// Count total rows first
	if err := r.db.Model(&entity.Event{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply offset and limit
	err := r.db.Offset(offset).Limit(limit).Find(&events).Error
	if err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

func (r *EventsRepository) Update(event *entity.Event) error {
	return r.db.Save(event).Error
}

func (r *EventsRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Event{}, id).Error
}
