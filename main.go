package main

import (
	"log"
	"lokajatim/config"
	articleController "lokajatim/controllers/article"
	authController "lokajatim/controllers/auth"
	commentController "lokajatim/controllers/comment"
	likeController "lokajatim/controllers/like"
	categoryController "lokajatim/controllers/category"
	productController "lokajatim/controllers/product"
	cartController "lokajatim/controllers/cart"
	TransactionController "lokajatim/controllers/transaction"
	PaymentController "lokajatim/controllers/payment"
	"lokajatim/controllers/event"
	"lokajatim/controllers/ticket"
	"lokajatim/middleware"
	articleRepo "lokajatim/repositories/article"
	authRepo "lokajatim/repositories/auth"
	commentRepo "lokajatim/repositories/comment"
	likeRepo "lokajatim/repositories/like"
	eventRepo "lokajatim/repositories/event"
	categoryRepo "lokajatim/repositories/category"
	productRepo "lokajatim/repositories/product"
	cartRepo "lokajatim/repositories/cart"
	TransactionRepo "lokajatim/repositories/transaction"
	PaymentRepo "lokajatim/repositories/payment"
	"lokajatim/routes"
	articleService "lokajatim/services/article"
	authService "lokajatim/services/auth"
	commentService "lokajatim/services/comment"
	likeService "lokajatim/services/like"
	eventService "lokajatim/services/event"
	ticketRepo "lokajatim/repositories/ticket"
	ticketService "lokajatim/services/ticket"
	categoryService "lokajatim/services/category"
	productService "lokajatim/services/product"
	cartService "lokajatim/services/cart"
	TransactionService "lokajatim/services/transaction"
	PaymentService "lokajatim/services/payment"

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

	// Initialize Echo server
	e := echo.New()

	// Initialize CORS middleware
	middleware.InitCors(e)

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
	transactionService := TransactionService.NewTransactionService(transactionRepo)
	transactionController := TransactionController.NewTransactionController(transactionService)

	// Initialize Payment components
	paymentRepo := PaymentRepo.NewPaymentRepository(db)
	paymentService := PaymentService.NewPaymentService(paymentRepo, transactionRepo)
	paymentController := PaymentController.NewPaymentController(paymentService)

	// Initialize RouteController with all controllers
	routeController := routes.RouteController{
		AuthController:    authController,
		ArticleController: articleController,
		CommentController: commentController,
		LikeController:    likeController,
		EventController:   eventController,
		TicketController:  ticketController,
		CategoryController: categoryController,
		ProductController: productController,
		CartController: cartController,
		TransactionController: transactionController,
		PaymentController: paymentController,
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
