package routes

import (
	"os"

	"lokajatim/controllers/auth"
	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	// Authentication routes
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	// Protected routes with JWT
	eJWT := e.Group("")
	eJWT.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
	}))
}