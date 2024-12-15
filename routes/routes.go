package routes

import (
	"net/http"
	"os"

	"lokajatim/controllers/article"
	"lokajatim/controllers/auth"
	"lokajatim/controllers/cart"
	"lokajatim/controllers/category"
	"lokajatim/controllers/chatbot"
	"lokajatim/controllers/comment"
	"lokajatim/controllers/event"
	"lokajatim/controllers/like"
	"lokajatim/controllers/product"
	"lokajatim/controllers/ticket"
	"lokajatim/controllers/transaction"

	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController        *auth.AuthController
	EventController       *event.EventController
	TicketController      *ticket.TicketController
	CommentController     *comment.CommentController
	ArticleController     *article.ArticleController
	LikeController        *like.LikeController
	CategoryController    *category.CategoryController
	ProductController     *product.ProductController
	CartController        *cart.CartController
	TransactionController *transaction.TransactionController
	ChatbotController     *chatbot.ChatbotController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	// Public Authentication Routes
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)
	// Route untuk forgot password (mengirim OTP)
	e.POST("/forgot-password", rc.AuthController.SendOTPController)

	// Route untuk reset password
	e.POST("/reset-password", rc.AuthController.ResetPasswordController)

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

	// Auth Routes
	eJWT.GET("/users", rc.AuthController.GetAllUsersController)
	eJWT.GET("/users/:id", rc.AuthController.GetUserByID)
	eJWT.PUT("/users/:id", rc.AuthController.UpdateUserController)
	eJWT.DELETE("/users/:id", rc.AuthController.DeleteUserController)

	// Event Routes
	e.GET("/events", rc.EventController.GetAllEvents)
	eJWT.GET("/events/:id", rc.EventController.GetEventByID)
	eJWT.POST("/events", rc.EventController.CreateEvent)
	eJWT.PUT("/events/:id", rc.EventController.UpdateEvent)
	eJWT.DELETE("/events/:id", rc.EventController.DeleteEvent)
	e.GET("/events/best", rc.EventController.GetByBestPrice)

	// Ticket Routes
	eJWT.GET("/tickets", rc.TicketController.GetAllTickets)
	eJWT.GET("/tickets/:id", rc.TicketController.GetTicketByID)
	eJWT.POST("/tickets", rc.TicketController.CreateTicket)
	eJWT.PUT("/tickets/:id", rc.TicketController.UpdateTicket)
	eJWT.DELETE("/tickets/:id", rc.TicketController.DeleteTicket)

	// Article Routes
	e.GET("/articles", rc.ArticleController.GetAll)
	e.GET("/articles/:id", rc.ArticleController.GetByID)
	e.POST("/articles", rc.ArticleController.Create)
	e.PUT("/articles/:id", rc.ArticleController.Update)
	e.DELETE("/articles/:id", rc.ArticleController.Delete)

	// Comment Routes
	e.GET("/comments/article/:article_id", rc.CommentController.GetCommentsByArticleID)
	e.GET("/comments/:id", rc.CommentController.GetCommentByID)
	e.POST("/comments", rc.CommentController.Create)
	e.DELETE("/comments/:id", rc.CommentController.Delete)

	// Like Routes
	e.POST("/likes", rc.LikeController.LikeArticle)
	e.DELETE("/likes/:article_id/:user_id", rc.LikeController.UnlikeArticle)
	e.GET("/likes/articles/:article_id", rc.LikeController.GetLikesByArticle)
	e.GET("/likes/articles/:article_id/count", rc.LikeController.CountLikes)
	e.GET("/likes/articles/:article_id/users/:user_id/status", rc.LikeController.IsUserLikedArticle)

	// Category Routes
	eJWT.GET("/categories", rc.CategoryController.GetCategories)
	eJWT.GET("/categories/:id", rc.CategoryController.GetCategoryByID)
	eJWT.POST("/categories", rc.CategoryController.CreateCategory)
	eJWT.PUT("/categories/:id", rc.CategoryController.UpdateCategory)
	eJWT.DELETE("/categories/:id", rc.CategoryController.DeleteCategory)

	// Product Routes
	e.GET("/products", rc.ProductController.GetAllProducts)
	e.GET("/products/:id", rc.ProductController.GetProductByID)
	eJWT.POST("/products", rc.ProductController.CreateProduct)
	eJWT.PUT("/products/:id", rc.ProductController.UpdateProduct)
	eJWT.DELETE("/products/:id", rc.ProductController.DeleteProduct)
	e.GET("/products/best", rc.ProductController.GetBestProductsPrice)
	eJWT.POST("/products/import", rc.ProductController.ImportProducts)

	// Cart Routes
	eJWT.GET("/carts/:user_id", rc.CartController.GetCartByUserID)
	eJWT.POST("/carts", rc.CartController.AddItemToCart)
	eJWT.PUT("/carts/:cart_item_id", rc.CartController.UpdateItemQuantity)
	eJWT.DELETE("/carts/:cart_id/clear", rc.CartController.ClearCart)
	eJWT.DELETE("/carts/:cart_item_id", rc.CartController.RemoveItemFromCart)

	// Transaction Routes
	eJWT.GET("/transactions", rc.TransactionController.GetAllTransactions)
	eJWT.GET("/transactions/:id", rc.TransactionController.GetTransactionByID)
	eJWT.POST("/transactions", rc.TransactionController.CreateTransaction)
	eJWT.PUT("/transactions/:id", rc.TransactionController.UpdateTransaction)
	eJWT.PUT("/transactions/:id/status", rc.TransactionController.UpdateTransactionStatus)
	eJWT.DELETE("/transactions/:id", rc.TransactionController.DeleteTransaction)
	e.POST("/transactions/notifications", rc.TransactionController.HandleMidtransNotification)

	// Chatbot Routes
	eJWT.POST("/chatbot", rc.ChatbotController.ChatbotController)
}
