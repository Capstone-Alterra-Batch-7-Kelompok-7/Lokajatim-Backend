package request

import "lokajatim/entities"

// LikeRequest is the request for the Like endpoint
// @Description LikeRequest is the request for the Like endpoint
// @Param ArticleID int true "Article ID"
// @Param UserID int true "User ID"
type LikeRequest struct {
	ArticleID 	int 	`json:"article_id"`
	UserID  	int    	`json:"user_id"`
}

func (likeRequest LikeRequest) ToEntities() (entities.Like, error) {
	return entities.Like{
		ArticleID: likeRequest.ArticleID,
		UserID: likeRequest.UserID,
	}, nil
}