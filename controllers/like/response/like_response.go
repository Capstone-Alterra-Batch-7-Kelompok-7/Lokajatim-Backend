package response

import "lokajatim/entities"

type LikeResponse struct {
	ID        int              `json:"id"`
	ArticleID int              `json:"article_id"`
	Article   entities.Article `json:"Article"`
	UserID    int              `json:"user_id"`
	User      entities.User    `json:"User"`
}

type CountLikesResponse struct {
	ArticleID int              `json:"article_id"`
	Count     int              `json:"count"`
}

type IsLikedResponse struct {
	ArticleID int              `json:"article_id"`
	UserID    int              `json:"user_id"`
	IsLiked   bool             `json:"is_liked"`
}

func FromLikeEntity(like entities.Like) LikeResponse {
	return LikeResponse{
		ID:        like.ID,
		ArticleID: like.ArticleID,
		Article:   like.Article,
		UserID:    like.UserID,
		User:      like.User,
	}
}
