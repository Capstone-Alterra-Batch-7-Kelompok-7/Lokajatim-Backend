package article

import (
	"lokajatim/controllers/base"
	"lokajatim/entities"
	"lokajatim/services/article"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
    ArticleService article.ArticleService
}

func NewArticleHandler(service article.ArticleService) *ArticleController {
    return &ArticleController{ArticleService: service}
}

func (h *ArticleController) GetAll(c echo.Context) error {
	articles, err := h.ArticleService.GetAllArticles()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, articles)
}

func (h *ArticleController) GetByID(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    article, err := h.ArticleService.GetArticleByID(uint(id))
    if err != nil {
        return base.ErrorResponse(c, err)
    }
    return base.SuccesResponse(c, article)
}

func (h *ArticleController) Create(c echo.Context) error {
	var article entities.Article
	if err := c.Bind(&article); err != nil {
		return base.ErrorResponse(c, err)
	}
	createdArticle, err := h.ArticleService.CreateArticle(&article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, createdArticle)
}

