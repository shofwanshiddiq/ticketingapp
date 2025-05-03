package main

import (
	"fmt"
	"net/http"
	"ticketingapp/config"
	"ticketingapp/entity"
	"ticketingapp/repositories"
	"ticketingapp/route"
	"ticketingapp/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}

	db := config.ConnectDatabase()

	db.AutoMigrate(
		&entity.User{},
		&entity.Event{},
		&entity.Ticket{},
		&entity.AuditLog{},
	)

	r := gin.Default()

	authRepository := repositories.NewAuthRepository(db)
	eventsReposity := repositories.NewEventsRepository(db)
	userRepository := repositories.NewUserRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	authService := services.NewAuthService(*authRepository)
	eventsService := services.NewEventsService(*eventsReposity)
	ticketService := services.NewTicketService(ticketRepository)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Shofwan Shiddiq's Ticketing App",
		})
	})

	route.SetupRoutes(r, authService, eventsService, userRepository, ticketService)

	// Start server
	port := ":8080"
	fmt.Printf("✅ Server is running on http://localhost%s\n", port)
	r.Run(port)

}
