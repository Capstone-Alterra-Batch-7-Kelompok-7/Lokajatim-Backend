package routes

import (
	"net/http"
	"os"

	"lokajatim/controllers/article"
	"lokajatim/controllers/auth"
	"lokajatim/controllers/comment"
	"lokajatim/controllers/like"

	"lokajatim/middleware"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController    *auth.AuthController
	CommentController *comment.CommentController
	ArticleController *article.ArticleController
	LikeController    *like.LikeController
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

	// Article Routes
	eJWT.GET("/articles", rc.ArticleController.GetAll)
	eJWT.GET("/articles/:id", rc.ArticleController.GetByID)
	eJWT.POST("/articles", rc.ArticleController.Create)
	eJWT.PUT("/articles/:id", rc.ArticleController.Update)
	eJWT.DELETE("/articles/:id", rc.ArticleController.Delete)

	// Comments Routes
	eJWT.GET("comments/articles/:article_id", rc.CommentController.GetCommentsByArticleID)
	eJWT.GET("/comments/:id", rc.CommentController.GetCommentByID)
	eJWT.POST("/comments", rc.CommentController.Create)
	eJWT.DELETE("/comments/:id", rc.CommentController.Delete)

	// Like Routes
	eJWT.POST("/likes", rc.LikeController.LikeArticle)
	eJWT.DELETE("/likes/:article_id/:user_id", rc.LikeController.UnlikeArticle)
	eJWT.GET("/likes/articles/:article_id", rc.LikeController.GetLikesByArticle)
	eJWT.GET("/likes/articles/:article_id/count", rc.LikeController.CountLikes)
	eJWT.GET("/likes/articles/:article_id/users/:user_id/status", rc.LikeController.IsUserLikedArticle)
}
