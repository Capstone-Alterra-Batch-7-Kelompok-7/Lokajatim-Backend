package routes

import (
	"net/http"
	"os"

	"lokajatim/controllers/auth"
	"lokajatim/controllers/event"
	"lokajatim/controllers/ticket"
	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	EventController *event.EventController
	TicketController *ticket.TicketController
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
	eJWT.GET("/events", rc.EventController.GetAllEvents)
	eJWT.GET("/events/:id", rc.EventController.GetEventByID)
	eJWT.POST("/events", rc.EventController.CreateEvent)
	eJWT.PUT("/events/:id", rc.EventController.UpdateEvent)
	eJWT.DELETE("/events/:id", rc.EventController.DeleteEvent)

	// Ticket routes
	eJWT.GET("/tickets", rc.TicketController.GetAllTickets)
	eJWT.GET("/tickets/:id", rc.TicketController.GetTicketByID)
	eJWT.POST("/tickets", rc.TicketController.CreateTicket)
	eJWT.PUT("/tickets/:id", rc.TicketController.UpdateTicket)
	eJWT.DELETE("/tickets/:id", rc.TicketController.DeleteTicket)
}
