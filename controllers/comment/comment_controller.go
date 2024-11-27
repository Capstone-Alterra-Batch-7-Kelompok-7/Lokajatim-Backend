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


func (h *CommentController) GetCommentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	comment, err := h.CommentService.GetCommentByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, comment)

}
func (h *CommentController) GetCommentsByArticleID(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	comments, err := h.CommentService.GetCommentsByArticleID(uint(articleID))
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
