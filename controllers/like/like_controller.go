package like

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/like/request"
	"lokajatim/controllers/like/response"
	"lokajatim/services/like"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeService like.LikeService
}

func NewLikeController(service like.LikeService) *LikeController {
	return &LikeController{likeService: service}
}

// @Summary Like an article
// @Tags Likes
// @Accept json
// @Produce json
// @Param request body request.LikeRequest true "Like Request"
// @Success 201 {object} response.LikeResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes [post]
func (c *LikeController) LikeArticle(ctx echo.Context) error {
	req := new(request.LikeRequest)
	if err := ctx.Bind(req); err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid request payload",
		})
	}

	likeEntity, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid request payload",
		})
	}

	created, err := c.likeService.LikeArticle(likeEntity.ArticleID, likeEntity.UserID)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to like article",
		})
	}

	return base.SuccesResponse(ctx, response.FromLikeEntity(created))
}

// @Summary Unlike an article
// @Tags Likes
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/{article_id}/{user_id} [delete]
func (c *LikeController) UnlikeArticle(ctx echo.Context) error {
	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid article ID",
		})
	}

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid user ID",
		})
	}

	err = c.likeService.UnlikeArticle(articleID, userID)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to unlike article",
		})
	}

	return base.SuccesResponse(ctx, map[string]string{"message": "unliked successfully"})
}

// @Summary Get all likes for an article
// @Tags Likes
// @Produce json
// @Success 200 {array} response.LikeResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/articles/{article_id} [get]
func (c *LikeController) GetLikesByArticle(ctx echo.Context) error {
	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid article ID",
		})
	}

	likes, err := c.likeService.GetLikesByArticle(articleID)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to get likes",
		})
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
// @Success 200 {object} response.CountLikesResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/articles/{article_id}/count [get]
func (c *LikeController) CountLikes(ctx echo.Context) error {
	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid article ID",
		})
	}

	count, err := c.likeService.CountLikes(articleID)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to count likes",
		})
	}

	return base.SuccesResponse(ctx, response.CountLikesResponse{
		ArticleID: articleID,
		Count:     count,
	})
}

// @Summary Check if a user liked an article
// @Tags Likes
// @Produce json
// @Success 200 {object} response.IsLikedResponse
// @Failure 400 {object} base.BaseResponse
// @Router /likes/{article_id}/users/{user_id}/status [get]
func (c *LikeController) IsUserLikedArticle(ctx echo.Context) error {
	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid article ID",
		})
	}

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Invalid user ID",
		})
	}

	isLiked, err := c.likeService.IsUserLikedArticle(articleID, userID)
	if err != nil {
		return base.ErrorResponse(ctx, err, map[string]string{
			"error": "Failed to check if user liked article",
		})
	}

	return base.SuccesResponse(ctx, response.IsLikedResponse{
		ArticleID: articleID,
		UserID:    userID,
		IsLiked:   isLiked,
	})
}
