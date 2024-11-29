package routes

import (
	"net/http"
	"os"

	"lokajatim/controllers/auth"
	"lokajatim/controllers/event"
	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	EventController *event.EventController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	// Authentication routes
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, CORS with Clean Architecture!")
	})

	// Protected routes with JWT
	eJWT := e.Group("")
	eJWT.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
	}))

	// Event routes
	e.GET("/events", rc.EventController.GetAllEvents)
	e.GET("/events/:id", rc.EventController.GetEventByID)
	e.POST("/events", rc.EventController.CreateEvent)
	e.PUT("/events/:id", rc.EventController.UpdateEvent)
	e.DELETE("/events/:id", rc.EventController.DeleteEvent)
}
