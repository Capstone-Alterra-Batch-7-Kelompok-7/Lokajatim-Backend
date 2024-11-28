package like

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/repositories/like"
)

type LikeService struct {
	likeRepo like.LikeRepoInterface
}

type LikeServiceInterface interface {
    LikeArticle(articleID, userID int) (entities.Like, error)
    UnlikeArticle(articleID, userID int) error
    GetLikesByArticle(articleID int) ([]entities.Like, error)
    CountLikes(articleID int) (int, error)
    IsUserLikedArticle(articleID, userID int) (bool, error)
}

func NewLikeService(repo like.LikeRepoInterface) LikeServiceInterface {
	return &LikeService{likeRepo: repo}
}

func (s *LikeService) LikeArticle(articleID, userID int) (entities.Like, error) {
	liked, err := s.likeRepo.IsUserLiked(articleID, userID)
	if err != nil {
		return entities.Like{}, err
	}
	if liked {
		return entities.Like{}, errors.New("user already liked this article")
	}

	like := entities.Like{
		ArticleID: articleID,
		UserID:    userID,
	}
	return s.likeRepo.CreateLike(like)
}

func (s *LikeService) UnlikeArticle(articleID, userID int) error {
	liked, err := s.likeRepo.IsUserLiked(articleID, userID)
	if err != nil {
		return err
	}
	if !liked {
		return errors.New("user has not liked this article")
	}
	return s.likeRepo.DeleteLike(articleID, userID)
}

func (s *LikeService) GetLikesByArticle(articleID int) ([]entities.Like, error) {
	return s.likeRepo.GetLikesByArticleID(articleID)
}

func (s *LikeService) CountLikes(articleID int) (int, error) {
	return s.likeRepo.CountLikesByArticleID(articleID)
}

func (s *LikeService) IsUserLikedArticle(articleID, userID int) (bool, error) {
	return s.likeRepo.IsUserLiked(articleID, userID)
}
