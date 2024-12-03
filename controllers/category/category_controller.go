package category

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/category/request"
	"lokajatim/controllers/category/response"
	"lokajatim/controllers/pagination"
	"lokajatim/services/category"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService category.CategoryService
}

func NewCategoryController(service category.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: service}
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Param request body request.CategoryRequest true "Category Request"
// @Success 200 {object} response.CategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /categories [get]
func (c *CategoryController) GetCategories(ctx echo.Context) error {
	categories, err := c.CategoryService.GetCategories()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to get categories",
		})
	}
	return pagination.SuccessPaginatedResponse(ctx, categories, 1, 10, int64(len(categories)))
}

// @Summary Get category by ID
// @Description Get category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the category"
// @Success 200 {object} response.CategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /categories/{id} [get]
func (c *CategoryController) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.CategoryService.GetCategoryByID(id)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to get category",
		})
	}
	return base.SuccesResponse(ctx, category)
}

// @Summary Create category
// @Description Create category
// @Tags Category
// @Accept json
// @Produce json
// @Param request body request.CategoryRequest true "Category Request"
// @Success 201 {object} response.CategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /categories [post]
func (c *CategoryController) CreateCategory(ctx echo.Context) error {
	req := new(request.CategoryRequest)
	if err := ctx.Bind(req); err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to bind request",
			})
			}
	category, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}
	created, err := c.CategoryService.CreateCategory(category)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to create category",
		})
	}
	return base.SuccesResponse(ctx, created)
}

// @Summary Update category
// @Description Update category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the category"
// @Param request body request.CategoryRequest true "Category Request"
// @Success 200 {object} response.CategoryResponse
// @Failure 400 {object} base.BaseResponse
// @Router /categories/{id} [put]
func (c *CategoryController) UpdateCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	req := new(request.CategoryRequest)
	if err := ctx.Bind(req); err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to bind request",
		})
	}
	category, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}
	updated, err := c.CategoryService.UpdateCategory(id, category)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to update category",
		})
	}
	return base.SuccesResponse(ctx, response.CategoryFromEntities(updated))
}

// @Summary Delete category
// @Description Delete category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "ID of the category"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /categories/{id} [delete]
func (c *CategoryController) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.CategoryService.DeleteCategory(id)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to delete category",
		})
	}
	return base.SuccesResponse(ctx, nil)
}