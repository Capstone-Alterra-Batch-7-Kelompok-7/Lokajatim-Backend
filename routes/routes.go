package routes

import (
	"net/http"
	"os"

	"lokajatim/controllers/article"
	"lokajatim/controllers/auth"
	"lokajatim/controllers/comment"
	"lokajatim/controllers/like"
	"lokajatim/controllers/event"
	"lokajatim/controllers/ticket"
	"lokajatim/controllers/category"
	"lokajatim/controllers/product"
	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController      *auth.AuthController
	EventController     *event.EventController
	TicketController    *ticket.TicketController
	CommentController   *comment.CommentController
	ArticleController   *article.ArticleController
	LikeController      *like.LikeController
	CategoryController  *category.CategoryController
	ProductController   *product.ProductController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	// Public Authentication Routes
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	// Basic route to check if API is up and running
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, CORS with Clean Architecture!")
	})

	// Grouped routes with JWT authentication
	eJWT := e.Group("")
	eJWT.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JwtCustomClaims)
		},
	}))

	// Event Routes
	eJWT.GET("/events", rc.EventController.GetAllEvents)
	eJWT.GET("/events/:id", rc.EventController.GetEventByID)
	eJWT.POST("/events", rc.EventController.CreateEvent)
	eJWT.PUT("/events/:id", rc.EventController.UpdateEvent)
	eJWT.DELETE("/events/:id", rc.EventController.DeleteEvent)

	// Ticket Routes
	eJWT.GET("/tickets", rc.TicketController.GetAllTickets)
	eJWT.GET("/tickets/:id", rc.TicketController.GetTicketByID)
	eJWT.POST("/tickets", rc.TicketController.CreateTicket)
	eJWT.PUT("/tickets/:id", rc.TicketController.UpdateTicket)
	eJWT.DELETE("/tickets/:id", rc.TicketController.DeleteTicket)

	// Article Routes
	eJWT.GET("/articles", rc.ArticleController.GetAll)
	eJWT.GET("/articles/:id", rc.ArticleController.GetByID)
	eJWT.POST("/articles", rc.ArticleController.Create)
	eJWT.PUT("/articles/:id", rc.ArticleController.Update)
	eJWT.DELETE("/articles/:id", rc.ArticleController.Delete)

	// Comment Routes
	eJWT.GET("/comments/article/:article_id", rc.CommentController.GetCommentsByArticleID)
	eJWT.GET("/comments/:id", rc.CommentController.GetCommentByID)
	eJWT.POST("/comments", rc.CommentController.Create)
	eJWT.DELETE("/comments/:id", rc.CommentController.Delete)

	// Like Routes
	eJWT.POST("/likes", rc.LikeController.LikeArticle)
	eJWT.DELETE("/likes/:article_id/:user_id", rc.LikeController.UnlikeArticle)
	eJWT.GET("/likes/articles/:article_id", rc.LikeController.GetLikesByArticle)
	eJWT.GET("/likes/articles/:article_id/count", rc.LikeController.CountLikes)
	eJWT.GET("/likes/articles/:article_id/users/:user_id/status", rc.LikeController.IsUserLikedArticle)

	// Category Routes
	eJWT.GET("/categories", rc.CategoryController.GetCategories)
	eJWT.GET("/categories/:id", rc.CategoryController.GetCategoryByID)
	eJWT.POST("/categories", rc.CategoryController.CreateCategory)
	eJWT.PUT("/categories/:id", rc.CategoryController.UpdateCategory)
	eJWT.DELETE("/categories/:id", rc.CategoryController.DeleteCategory)

	// Product Routes
	eJWT.GET("/products", rc.ProductController.GetAllProducts)
	eJWT.GET("/products/:id", rc.ProductController.GetProductByID)
	eJWT.POST("/products", rc.ProductController.CreateProduct)
	eJWT.PUT("/products/:id", rc.ProductController.UpdateProduct)
	eJWT.DELETE("/products/:id", rc.ProductController.DeleteProduct)
}
