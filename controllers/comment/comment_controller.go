package comment

import (
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


func (h *CommentController) GetCommentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	comment, err := h.CommentService.GetCommentByID((id))
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

	comments, err := h.CommentService.GetCommentsByArticleID(articleID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return pagination.SuccessPaginatedResponse(c, comments, 1, 10, int64(len(comments)))
}

func (h *CommentController) Create(c echo.Context) error {
    req := new(request.CommentRequest)
    if err := c.Bind(&req); err != nil {
        return base.ErrorResponse(c, err)
    }

	comment, err := req.ToEntities()
	if err != nil {
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
    if err := h.CommentService.DeleteComment(id); err != nil {
        return base.ErrorResponse(c, err)
    }
    return base.SuccesResponse(c, map[string]string{"message": "Comment deleted"})
}
