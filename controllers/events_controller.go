package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketingapp/entity"
	"ticketingapp/services"
	"ticketingapp/types"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	service services.EventsService
}

func NewEventsController(service services.EventsService) *EventsController {
	return &EventsController{service}
}

func (ec *EventsController) FindAll(c *gin.Context) {
	// Default values
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	// Convert to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	// Fetch data with pagination
	events, totalItems, err := ec.service.FindAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format response using response type mapping
	eventResponses := types.ToEventResponseList(events)

	// Calculate total pages
	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	// Return paginated response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    eventResponses,
		"page":    page,
		"limit":   limit,
		"total":   totalItems,
		"pages":   totalPages,
		"message": "Events fetched successfully",
	})
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

	event, err := ec.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := ec.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Event with ID %d deleted successfully", event.ID),
	})
}
