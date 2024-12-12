package response

import "lokajatim/entities"

// LikeResponse is the response for the Like endpoint
// @Description LikeResponse is the response for the Like endpoint
// @Param ID int true "ID of the like"
// @Param ArticleID int true "ID of the article"
// @Param Article Article true "Article of the like"
// @Param UserID int true "ID of the user"
// @Param User User true "User of the like"
type LikeResponse struct {
	ID        int              `json:"id"`
	ArticleID int              `json:"article_id"`
	Article   entities.Article `json:"Article"`
	UserID    int              `json:"user_id"`
	User      entities.User    `json:"User"`
}

// CountLikesResponse is the response for the CountLikes endpoint
// @Description CountLikesResponse is the response for the CountLikes endpoint
// @Param ArticleID int true "ID of the article"
// @Param Count int true "Count of the likes"
type CountLikesResponse struct {
	ArticleID int              `json:"article_id"`
	Count     int              `json:"count"`
}

// IsLikedResponse is the response for the IsLiked endpoint
// @Description IsLikedResponse is the response for the IsLiked endpoint
// @Param ArticleID int true "ID of the article"
// @Param UserID int true "ID of the user"
// @Param IsLiked bool true "Is liked or not"
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
