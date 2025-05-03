package route

import (
	"ticketingapp/controllers"
	"ticketingapp/middleware"
	"ticketingapp/repositories"
	"ticketingapp/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authService services.AuthService, eventRepo services.EventsService, userRepo repositories.UserRepository, ticketRepo services.TicketService) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			authController := controllers.NewAuthController(authService)
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(userRepo))
	{
		events := protected.Group("/events")
		{
			eventsController := controllers.NewEventsController(eventRepo)
			events.GET("/", eventsController.FindAll)
			events.GET("/:id", eventsController.FindByID)
			events.POST("/", eventsController.Create)
			events.PUT("/:id", eventsController.Update)
			events.DELETE("/:id", eventsController.Delete)
		}

		tickets := protected.Group("/tickets")
		{
			ticketController := controllers.NewTicketController(ticketRepo)
			tickets.GET("/", ticketController.GetAllTickets)
			tickets.GET("/:id", ticketController.GetTicketByID)
			tickets.POST("/", ticketController.CreateTicket)
			tickets.PATCH("/:id", ticketController.CancelTicket)
		}
	}
}
