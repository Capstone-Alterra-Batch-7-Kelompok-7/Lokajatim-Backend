package like

import (
	"errors"
	"lokajatim/entities"

	"gorm.io/gorm"
)

type LikeRepo struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepoInterface {
	return &LikeRepo{db: db}
}

func (r *LikeRepo) CreateLike(like entities.Like) (entities.Like, error) {
	if err := r.db.Create(&like).Error; err != nil {
		return entities.Like{}, err
	}

	var createdLike entities.Like
    result := r.db.Preload("Article").Preload("User").First(&createdLike, like.ID)
    if result.Error != nil {
        return entities.Like{}, result.Error
    }
    return createdLike, nil
}

func (r *LikeRepo) DeleteLike(articleID, userID int) error {
	if err := r.db.Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&entities.Like{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *LikeRepo) GetLikesByArticleID(articleID int) ([]entities.Like, error) {
	var likes []entities.Like
	if err := r.db.Preload("Article").Preload("User").Where("article_id = ?", articleID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (r *LikeRepo) CountLikesByArticleID(articleID int) (int, error) {
	var count int64
	if err := r.db.Preload("Article").Model(&entities.Like{}).Where("article_id = ?", articleID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *LikeRepo) IsUserLiked(articleID, userID int) (bool, error) {
	var like entities.Like
	if err := r.db.Preload("Article").Preload("User").Where("article_id = ? AND user_id = ?", articleID, userID).First(&like).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
