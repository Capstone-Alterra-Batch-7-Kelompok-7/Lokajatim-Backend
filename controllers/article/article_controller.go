package article

import (
	"lokajatim/controllers/article/request"
	"lokajatim/controllers/base"
	"lokajatim/controllers/pagination"
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
	return pagination.SuccessPaginatedResponse(c, articles, 1, 10, int64(len(articles)))
}

func (h *ArticleController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleService.GetArticleByID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, article)
}

func (h *ArticleController) Create(c echo.Context) error {
	req := new(request.ArticleRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}
	article, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	created, err := h.ArticleService.CreateArticle(&article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, created)
}

func (h *ArticleController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article := new(entities.Article)
	if err := c.Bind(article); err != nil {
		return base.ErrorResponse(c, err)
	}
	updatedArticle, err := h.ArticleService.UpdateArticle(id, article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, updatedArticle)
}

func (h *ArticleController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ArticleService.DeleteArticle(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "Article deleted successfully")
}
