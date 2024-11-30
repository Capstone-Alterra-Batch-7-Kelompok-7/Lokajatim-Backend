package request

import "lokajatim/entities"

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