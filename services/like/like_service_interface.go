package like

import "lokajatim/entities"

type LikeServiceInterface interface {
	LikeArticle(articleID, userID int) (entities.Like, error)
	UnlikeArticle(articleID, userID int) error
	GetLikesByArticle(articleID int) ([]entities.Like, error)
	CountLikes(articleID int) (int, error)
	IsUserLikedArticle(articleID, userID int) (bool, error)
}