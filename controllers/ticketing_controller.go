package controllers

import (
	"net/http"
	"strconv"

	"ticketingapp/entity"
	"ticketingapp/services"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	service services.TicketService
}

func NewTicketController(service services.TicketService) *TicketController {
	return &TicketController{service: service}
}

func (c *TicketController) GetAllTickets(ctx *gin.Context) {
	tickets, err := c.service.GetAllTickets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

func (c *TicketController) GetTicketByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket ID"})
		return
	}

	ticket, err := c.service.GetTicketByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ticket)
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var ticket entity.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateTicket(&ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, ticket)
}

func (c *TicketController) CancelTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket ID"})
		return
	}

	if err := c.service.CancelTicket(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ticket successfully cancelled"})
}
