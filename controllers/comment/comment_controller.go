package comment

import (
	"lokajatim/controllers/base"
	"lokajatim/entities"
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

func (h *CommentController) GetAllComments(c echo.Context) error {
	comments, err := h.CommentService.GetAllComments()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, comments)
}

func (h *CommentController) Create(c echo.Context) error {
    var comment entities.Comment
    if err := c.Bind(&comment); err != nil {
        return base.ErrorResponse(c, err)
    }
    created, err := h.CommentService.CreateComment(comment)
    if err != nil {
        return base.ErrorResponse(c, err)
    }
    return base.SuccesResponse(c, created)
}

func (h *CommentController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.CommentService.DeleteComment(uint(id)); err != nil {
        return base.ErrorResponse(c, err)
    }
    return base.SuccesResponse(c, map[string]string{"message": "Comment deleted"})
}
