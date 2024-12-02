package main

import (
	"log"
	"lokajatim/config"
	articleController "lokajatim/controllers/article"
	authController "lokajatim/controllers/auth"
	commentController "lokajatim/controllers/comment"
	likeController "lokajatim/controllers/like"
	"lokajatim/controllers/event"
	"lokajatim/controllers/ticket"
	"lokajatim/middleware"
	articleRepo "lokajatim/repositories/article"
	authRepo "lokajatim/repositories/auth"
	commentRepo "lokajatim/repositories/comment"
	likeRepo "lokajatim/repositories/like"
	eventRepo "lokajatim/repositories/event"
	"lokajatim/routes"
	articleService "lokajatim/services/article"
	authService "lokajatim/services/auth"
	commentService "lokajatim/services/comment"
	likeService "lokajatim/services/like"
	eventService "lokajatim/services/event"
	ticketRepo "lokajatim/repositories/ticket"
	ticketService "lokajatim/services/ticket"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "lokajatim/docs"
)

// @title Lokajatim API
// @version 1.0
// @description This is the API documentation for Lokajatim.

// @host localhost:8000
// @BasePath /
func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Database
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	config.MigrateDB(db)

	// Initialize Echo
	e := echo.New()

	// Initialize CORS middleware
	middleware.InitCors(e)

	// Initialize Auth
	authJwt := middleware.JwtLokajatim{}
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	// Initialize Article components
	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(*articleService)

	// Initialize Comment components
	commentRepo := commentRepo.NewCommentRepository(db)
	commentService := commentService.NewCommentService(commentRepo)
	commentController := commentController.NewCommentController(*commentService)

	// Initialize Like components
	likeRepo := likeRepo.NewLikeRepository(db)
	likeService := likeService.NewLikeService(likeRepo)
	likeController := likeController.NewLikeController(*likeService)

	// Initialize Event
	eventRepo := eventRepo.NewEventRepo(db)
	eventService := eventService.NewEventService(eventRepo)
	eventController := event.NewEventController(eventService)

	// Initialize Ticket
	ticketRepo := ticketRepo.NewTicketRepository(db)
	ticketService := ticketService.NewTicketService(ticketRepo)
	ticketController := ticket.NewTicketController(ticketService, authService, eventService)	

	// Initialize RouteController with all controllers
	routeController := routes.RouteController{
		AuthController:    authController,
		ArticleController: articleController,
		CommentController: commentController,
		LikeController:    likeController,
		EventController:  eventController, 
		TicketController: ticketController,
	}
	
	// Setup routes
	routeController.InitRoute(e)

	// CORS middleware
	middleware.InitCors(e)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start the Echo server
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// loadEnv loads the environment variables from a .env file
func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
