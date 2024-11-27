package main

import (
	"log"
	"lokajatim/config"
	authController "lokajatim/controllers/auth"
	articleController "lokajatim/controllers/article"
	commentController "lokajatim/controllers/comment"
	"lokajatim/middleware"
	authRepo "lokajatim/repositories/auth"
	articleRepo "lokajatim/repositories/article"
	commentRepo "lokajatim/repositories/comment"
	"lokajatim/routes"
	authService "lokajatim/services/auth"
	articleService "lokajatim/services/article"
	commentService "lokajatim/services/comment"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	config.MigrateDB(db)

	// Initialize Echo
	e := echo.New()

	// Initialize Auth components
	authJwt := middleware.JwtLokajatim{}
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	// Initialize Article components
	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(articleService)

	// Initialize Comment components
	commentRepo := commentRepo.NewCommentRepository(db)
	commentService := commentService.NewCommentService(commentRepo)
	commentController := commentController.NewCommentController(*commentService)

	// Initialize RouteController with all controllers
	routeController := routes.RouteController{
		AuthController:    authController,
		ArticleController: articleController,
		CommentController: commentController,
	}

	// Setup routes
	routeController.InitRoute(e)

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
