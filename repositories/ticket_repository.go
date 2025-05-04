package repositories

import (
	"errors"
	"sync"

	"ticketingapp/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAllTickets() ([]entity.Ticket, error)
	GetTicketByID(id uint) (*entity.Ticket, error)
	CreateTicket(ticket *entity.Ticket) error
	UpdateTicketStatus(id uint, status string) error
}

type ticketRepository struct {
	db *gorm.DB
	mu sync.Mutex
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) GetAllTickets() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	if err := r.db.Preload("User").Preload("Event").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(id uint) (*entity.Ticket, error) {
	var ticket entity.Ticket
	if err := r.db.Preload("User").Preload("Event").First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) CreateTicket(ticket *entity.Ticket) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check event capacity
	var event entity.Event
	if err := r.db.First(&event, ticket.EventID).Error; err != nil {
		return err
	}
	if event.Capacity <= 0 {
		return errors.New("event capacity is full")
	}

	if err := r.db.Create(ticket).Error; err != nil {
		return err
	}

	event.Capacity -= 1
	if err := r.db.Save(&event).Error; err != nil {
		return err
	}

	return nil
}

func (r *ticketRepository) UpdateTicketStatus(id uint, status string) error {
	var ticket entity.Ticket
	if err := r.db.First(&ticket, id).Error; err != nil {
		return err
	}

	// Update status
	ticket.Status = status
	if err := r.db.Save(&ticket).Error; err != nil {
		return err
	}
	return nil
}
