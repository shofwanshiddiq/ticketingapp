package main

import (
	"fmt"
	"ticketingapp/config"
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

	r := gin.Default()

	config.RegisterRootRoute(r)

	authRepository := repositories.NewAuthRepository(db)
	eventsReposity := repositories.NewEventsRepository(db)
	userRepository := repositories.NewUserRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	authService := services.NewAuthService(*authRepository)
	eventsService := services.NewEventsService(*eventsReposity)
	ticketService := services.NewTicketService(ticketRepository)

	route.SetupRoutes(r, authService, eventsService, userRepository, ticketService)

	// Start server
	port := ":8080"
	fmt.Printf("✅ Server is running on http://localhost%s\n", port)
	r.Run(port)
}
