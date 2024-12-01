package main

import (
	"log"
	"lokajatim/config"
	authController "lokajatim/controllers/auth"
	"lokajatim/controllers/event"
	"lokajatim/controllers/ticket"
	"lokajatim/middleware"
	authRepo "lokajatim/repositories/auth"
	eventRepo "lokajatim/repositories/event"
	"lokajatim/routes"
	authService "lokajatim/services/auth"
	eventService "lokajatim/services/event"
	ticketRepo "lokajatim/repositories/ticket"
	ticketService "lokajatim/services/ticket"
	_ "lokajatim/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Lokajatim API
// @version 1.0
// @description This is the API documentation for Lokajatim.

// @host localhost:8000
// @BasePath /
func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Database
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	config.MigrateDB(db)

	e := echo.New()

	// Initialize CORS middleware
	middleware.InitCors(e)

	// Initialize Auth
	authJwt := middleware.JwtLokajatim{}
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	// Initialize Event
	eventRepo := eventRepo.NewEventRepo(db)
	eventService := eventService.NewEventService(eventRepo)
	eventController := event.NewEventController(eventService)

	// Initialize Ticket
	ticketRepo := ticketRepo.NewTicketRepository(db)
	ticketService := ticketService.NewTicketService(ticketRepo)
	ticketController := ticket.NewTicketController(ticketService, authService, eventService)	

	// Initialize Routes
	routeController := routes.RouteController{
		AuthController:   authController,
		EventController:  eventController, 
		TicketController: ticketController,
	}
	routeController.InitRoute(e)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
