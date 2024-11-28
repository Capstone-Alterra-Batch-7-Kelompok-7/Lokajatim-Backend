package like

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/like/response"
	"lokajatim/services/like"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeService like.LikeServiceInterface
}

func NewLikeController(service like.LikeServiceInterface) *LikeController {
	return &LikeController{likeService: service}
}

// @Summary Like an article
// @Tags Likes
// @Accept json
// @Produce json
// @Param article_id path int true "Article ID"
// @Param user_id path int true "User ID"
// @Success 201 {object} entities.Like
// @Failure 400 {object} base.BaseResponse
// @Router /likes [post]
func (c *LikeController) LikeArticle(ctx echo.Context) error {
	articleID, _ := strconv.Atoi(ctx.Param("article_id"))
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	like, err := c.likeService.LikeArticle(articleID, userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, like)
}

// @Summary Unlike an article
// @Tags Likes
// @Param article_id path int true "Article ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes [delete]
func (c *LikeController) UnlikeArticle(ctx echo.Context) error {
	articleID, _ := strconv.Atoi(ctx.Param("article_id"))
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	err := c.likeService.UnlikeArticle(articleID, userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "unliked successfully"})
}


// @Summary Get all likes for an article
// @Tags Likes
// @Produce json
// @Param article_id path int true "Article ID"
// @Success 200 {array} response.LikeResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/{article_id} [get]
func (c *LikeController) GetLikesByArticle(ctx echo.Context) error {
	articleID, _ := strconv.Atoi(ctx.Param("article_id"))

	likes, err := c.likeService.GetLikesByArticle(articleID)
	if err != nil {
		return base.ErrorResponse(ctx, err)
	}

	likeResponses := make([]response.LikeResponse, len(likes))
	for i, like := range likes {
		likeResponses[i] = response.FromLikeEntity(like)
	}

	return base.SuccesResponse(ctx, likeResponses)
}

// @Summary Count likes for an article
// @Tags Likes
// @Produce json
// @Param article_id path int true "Article ID"
// @Success 200 {object} response.CountLikesResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/count/{article_id} [get]
func (c *LikeController) CountLikes(ctx echo.Context) error {
	articleID, _ := strconv.Atoi(ctx.Param("article_id"))

	count, err := c.likeService.CountLikes(articleID)
	if err != nil {
		return base.ErrorResponse(ctx, err)
	}

	return base.SuccesResponse(ctx, response.CountLikesResponse{
		ArticleID: articleID,
		Count:     count,
	})
}

// @Summary Check if a user liked an article
// @Tags Likes
// @Produce json
// @Param article_id path int true "Article ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} response.IsLikedResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/{article_id}/{user_id}/status [get]
func (c *LikeController) IsUserLikedArticle(ctx echo.Context) error {
	articleID, _ := strconv.Atoi(ctx.Param("article_id"))
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	isLiked, err := c.likeService.IsUserLikedArticle(articleID, userID)
	if err != nil {
		return base.ErrorResponse(ctx, err)
	}

	return base.SuccesResponse(ctx, response.IsLikedResponse{
		ArticleID: articleID,
		UserID:    userID,
		IsLiked:   isLiked,
	})
}
