package main

import (
	"log"
	"lokajatim/config"
	articleController "lokajatim/controllers/article"
	authController "lokajatim/controllers/auth"
	cartController "lokajatim/controllers/cart"
	categoryController "lokajatim/controllers/category"
	commentController "lokajatim/controllers/comment"
	"lokajatim/controllers/event"
	likeController "lokajatim/controllers/like"
	productController "lokajatim/controllers/product"
	"lokajatim/controllers/ticket"
	TransactionController "lokajatim/controllers/transaction"
	"lokajatim/middleware"
	articleRepo "lokajatim/repositories/article"
	authRepo "lokajatim/repositories/auth"
	cartRepo "lokajatim/repositories/cart"
	categoryRepo "lokajatim/repositories/category"
	commentRepo "lokajatim/repositories/comment"
	eventRepo "lokajatim/repositories/event"
	likeRepo "lokajatim/repositories/like"
	productRepo "lokajatim/repositories/product"
	ticketRepo "lokajatim/repositories/ticket"
	TransactionRepo "lokajatim/repositories/transaction"
	"lokajatim/routes"
	articleService "lokajatim/services/article"
	authService "lokajatim/services/auth"
	cartService "lokajatim/services/cart"
	categoryService "lokajatim/services/category"
	commentService "lokajatim/services/comment"
	eventService "lokajatim/services/event"
	likeService "lokajatim/services/like"
	productService "lokajatim/services/product"
	ticketService "lokajatim/services/ticket"
	TransactionService "lokajatim/services/transaction"
	"lokajatim/utils"

	_ "lokajatim/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	// Initialize Echo server
	e := echo.New()

	// Initialize CORS middleware
	middleware.InitCors(e)

	// Initialize Midtrans
	utils.InitMidtrans()

	// Initialize Auth components
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

	// Initialize Event components
	eventRepo := eventRepo.NewEventRepo(db)
	eventService := eventService.NewEventService(eventRepo)
	eventController := event.NewEventController(eventService)

	// Initialize Ticket components
	ticketRepo := ticketRepo.NewTicketRepository(db)
	ticketService := ticketService.NewTicketService(ticketRepo)
	ticketController := ticket.NewTicketController(ticketService, authService, eventService)

	// Initialize Category components
	categoryRepo := categoryRepo.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepo)
	categoryController := categoryController.NewCategoryController(*categoryService)

	// Initialize Product components
	productRepo := productRepo.NewProductRepository(db)
	productService := productService.NewProductService(productRepo)
	productController := productController.NewProductController(*productService)

	// Initialize Cart components
	cartRepo := cartRepo.NewCartRepository(db)
	cartService := cartService.NewCartService(cartRepo)
	cartController := cartController.NewCartController(*cartService)

	// Initialize Transaction components
	transactionRepo := TransactionRepo.NewTransactionRepository(db)
	transactionService := TransactionService.NewTransactionService(transactionRepo, cartRepo)
	transactionController := TransactionController.NewTransactionController(transactionService)

	// Initialize RouteController with all controllers
	routeController := routes.RouteController{
		AuthController:        authController,
		ArticleController:     articleController,
		CommentController:     commentController,
		LikeController:        likeController,
		EventController:       eventController,
		TicketController:      ticketController,
		CategoryController:    categoryController,
		ProductController:     productController,
		CartController:        cartController,
		TransactionController: transactionController,
	}

	// Set up all routes using the routeController
	routeController.InitRoute(e)

	// Swagger documentation endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start Echo server on port 8000
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// loadEnv loads environment variables from a .env file
func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
