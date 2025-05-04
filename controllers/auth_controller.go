package controllers

import (
	"ticketingapp/dto"
	"ticketingapp/entity"
	"ticketingapp/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user entity.User
	// pake dto
	if err := c.ShouldBindJSON(&user); err != nil {
		// pake type
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := ac.service.Register(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := ac.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
