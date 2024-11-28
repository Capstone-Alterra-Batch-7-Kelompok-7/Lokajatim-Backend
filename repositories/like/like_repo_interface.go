package like

import "lokajatim/entities"

type LikeRepoInterface interface {
    CreateLike(like entities.Like) (entities.Like, error)
    DeleteLike(articleID, userID int) error
    GetLikesByArticleID(articleID int) ([]entities.Like, error)
    CountLikesByArticleID(articleID int) (int, error)
    IsUserLiked(articleID, userID int) (bool, error)
}
