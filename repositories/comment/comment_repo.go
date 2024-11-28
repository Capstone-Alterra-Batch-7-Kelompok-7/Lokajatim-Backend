package comment

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type commentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) GetCommentByID(id int) (entities.Comment, error) {
	var comment entities.Comment
	result := r.db.Preload("Article").First(&comment, id)
	if result.Error != nil {
		return entities.Comment{}, result.Error
	}
	return comment, nil
}

func (r *commentRepositoryImpl) GetCommentsByArticleID(articleID int) ([]entities.Comment, error) {
	var comments []entities.Comment
	result := r.db.Preload("Article").Where("article_id = ?", articleID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (r *commentRepositoryImpl) CreateComment(comment entities.Comment) (entities.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return entities.Comment{}, err
	}

	var createdComment entities.Comment
	result := r.db.Preload("Article").First(&createdComment, comment.ID)
	if result.Error != nil {
		return entities.Comment{}, result.Error
	}
	return createdComment, nil
}

func (r *commentRepositoryImpl) DeleteComment(id int) error {
	var comment entities.Comment
	if err := r.db.Delete(&comment, id).Error; err != nil {
		return err
	}
	return nil
}
