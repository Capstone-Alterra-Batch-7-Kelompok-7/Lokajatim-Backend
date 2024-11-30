package main

import (
	"log"
	"lokajatim/config"
	articleController "lokajatim/controllers/article"
	authController "lokajatim/controllers/auth"
	commentController "lokajatim/controllers/comment"
	likeController "lokajatim/controllers/like"
	"lokajatim/middleware"
	articleRepo "lokajatim/repositories/article"
	authRepo "lokajatim/repositories/auth"
	commentRepo "lokajatim/repositories/comment"
	likeRepo "lokajatim/repositories/like"
	"lokajatim/routes"
	articleService "lokajatim/services/article"
	authService "lokajatim/services/auth"
	commentService "lokajatim/services/comment"
	likeService "lokajatim/services/like"

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
	articleController := articleController.NewArticleController(*articleService)

	// Initialize Comment components
	commentRepo := commentRepo.NewCommentRepository(db)
	commentService := commentService.NewCommentService(commentRepo)
	commentController := commentController.NewCommentController(*commentService)

	// Initialize Like components
	likeRepo := likeRepo.NewLikeRepository(db)
	likeService := likeService.NewLikeService(likeRepo)
	likeController := likeController.NewLikeController(*likeService)

	// Initialize RouteController with all controllers
	routeController := routes.RouteController{
		AuthController:    authController,
		ArticleController: articleController,
		CommentController: commentController,
		LikeController:    likeController,
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
