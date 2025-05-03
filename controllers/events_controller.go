package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketingapp/entity"
	"ticketingapp/services"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	service services.EventsService
}

func NewEventsController(service services.EventsService) *EventsController {
	return &EventsController{service}
}

func (ec *EventsController) FindAll(c *gin.Context) {
	events, err := ec.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)
}

func (ec *EventsController) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	id := uint(idUint64)

	event, err := ec.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (ec *EventsController) Create(c *gin.Context) {
	var event entity.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := ec.service.Create(&event)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, event)
}

func (ec *EventsController) Update(c *gin.Context) {
	var event entity.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := ec.service.Update(&event)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, event)
}

func (ec *EventsController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Validasi: cek apakah event dengan ID tsb ada
	event, err := ec.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Lanjut delete
	if err := ec.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Event with ID %d deleted successfully", event.ID),
	})
}
