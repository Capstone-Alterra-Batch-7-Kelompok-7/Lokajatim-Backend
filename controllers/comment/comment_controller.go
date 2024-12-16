package comment

import (
	"lokajatim/controllers/comment/response"
	"lokajatim/controllers/base"
	"lokajatim/controllers/comment/request"
	"lokajatim/controllers/pagination"
	"lokajatim/services/comment"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
    CommentService comment.CommentService
}

func NewCommentController(service comment.CommentService) *CommentController {
    return &CommentController{CommentService: service}
}

// @Summary Get comment by ID
// @Description Get comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Success 200 {object} response.CommentResponse
// @Failure 400 {object} base.BaseResponse
// @Router /comments/{id} [get]
func (h *CommentController) GetCommentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid comment ID",
		})
	}

	comment, err := h.CommentService.GetCommentByID((id))
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get comment",
		})
	}

	return base.SuccesResponse(c, response.CommentFromEntities(comment))

}
// @Summary Get comments by article ID
// @Description Get comments by article ID
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} response.CommentResponse
// @Failure 400 {object} base.BaseResponse
// @Router /comments/article/{article_id} [get]
func (h *CommentController) GetCommentsByArticleID(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid article ID",
		})
	}

	comments, err := h.CommentService.GetCommentsByArticleID(articleID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get comments",
		})
	}
	return pagination.SuccessPaginatedResponse(c, comments, 1, 10, int64(len(comments)))
}

// @Summary Create comment
// @Description Create comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param request body request.CommentRequest true "Comment Request"
// @Success 201 {object} response.CommentResponse
// @Failure 400 {object} base.BaseResponse
// @Router /comments [post]
func (h *CommentController) Create(c echo.Context) error {
    req := new(request.CommentRequest)
    if err := c.Bind(&req); err != nil {
        return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
    }

	comment, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}

    created, err := h.CommentService.CreateComment(comment)
    if err != nil {
        return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to create comment",
		})
    }
    return base.SuccesResponse(c, created)
}

// @Summary Delete comment
// @Description Delete comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /comments/{id} [delete]
func (h *CommentController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.CommentService.DeleteComment(id); err != nil {
        return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to delete comment",
		})
    }
    return base.SuccesResponse(c, map[string]string{"message": "Comment deleted"})
}
