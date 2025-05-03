package services

import (
	"errors"
	"sync"
	"ticketingapp/entity"
	"ticketingapp/repositories"
)

type TicketService interface {
	GetAllTickets() ([]entity.Ticket, error)
	GetTicketByID(id uint) (*entity.Ticket, error)
	CreateTicket(ticket *entity.Ticket) error
	CancelTicket(id uint) error
}

type ticketService struct {
	repo repositories.TicketRepository
}

func NewTicketService(repo repositories.TicketRepository) TicketService {
	return &ticketService{repo: repo}
}

func (s *ticketService) GetAllTickets() ([]entity.Ticket, error) {
	return s.repo.GetAllTickets()
}

func (s *ticketService) GetTicketByID(id uint) (*entity.Ticket, error) {
	return s.repo.GetTicketByID(id)
}

func (s *ticketService) CreateTicket(ticket *entity.Ticket) error {
	if ticket.Type == "" {
		return errors.New("ticket type is required")
	}

	if ticket.Status != "tersedia" {
		return errors.New("invalid ticket status")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.repo.CreateTicket(ticket)
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	return nil
}

func (s *ticketService) CancelTicket(id uint) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil {
		return err
	}

	if ticket.Status == "dibatalkan" || ticket.Status == "habis" {
		return errors.New("cannot cancel ticket with current status")
	}

	return s.repo.UpdateTicketStatus(id, "dibatalkan")
}
