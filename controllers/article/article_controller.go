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

func NewArticleController(service article.ArticleService) *ArticleController {
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

func (h *ArticleController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var article entities.Article
	if err := c.Bind(&article); err != nil {
		return base.ErrorResponse(c, err)
	}
	updatedArticle, err := h.ArticleService.UpdateArticle(uint(id), &article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, updatedArticle)
}

func (h *ArticleController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ArticleService.DeleteArticle(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "Article deleted successfully")
}

