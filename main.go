package main

import (
    "log"
    "lokajatim/config"
    authController "lokajatim/controllers/auth"
    "lokajatim/middleware"
    authRepo "lokajatim/repositories/auth"
    "lokajatim/routes"
    authService "lokajatim/services/auth"

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

    db, err := config.InitDatabase()
    if err != nil {
        log.Fatalf("Database initialization failed: %v", err)
    }
    config.MigrateDB(db)

    e := echo.New()

    // Initialize Auth
    authJwt := middleware.JwtLokajatim{}
    authRepo := authRepo.NewAuthRepo(db)
    authService := authService.NewAuthService(authRepo, authJwt)
    authController := authController.NewAuthController(authService)

    routeController := routes.RouteController{
        AuthController: authController,
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
