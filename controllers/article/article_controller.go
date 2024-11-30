package article

import (
	"lokajatim/controllers/article/request"
	"lokajatim/controllers/article/response"
	"lokajatim/controllers/base"
	"lokajatim/controllers/pagination"
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
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get articles",
		})
	}
	return pagination.SuccessPaginatedResponse(c, articles, 1, 10, int64(len(articles)))
}

func (h *ArticleController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleService.GetArticleByID(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get article",
		})
	}
	return base.SuccesResponse(c, article)
}

func (h *ArticleController) Create(c echo.Context) error {
	req := new(request.ArticleRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}
	article, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}
	created, err := h.ArticleService.CreateArticle(article)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to create article",
		})
	}
	return base.SuccesResponse(c, created)
}

func (h *ArticleController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.ArticleRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}
	article, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}
	updated, err := h.ArticleService.UpdateArticle(id, article)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update article",
		})
	}
	return base.SuccesResponse(c, response.ArticleFromEntities(updated))
}

func (h *ArticleController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ArticleService.DeleteArticle(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to delete article",
		})
	}
	return base.SuccesResponse(c, "Article deleted successfully")
}
