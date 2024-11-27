package comment

import (
	"lokajatim/entities"
	"gorm.io/gorm"

)

type commentRepositoryImpl struct{
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
    return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) GetCommentByID(id uint) (entities.Comment, error) {
	var comment entities.Comment
	result := r.db.First(&comment, id)
	if result.Error != nil {
		return entities.Comment{}, result.Error
	}
	return comment, nil
}

func (r *commentRepositoryImpl) GetCommentsByArticleID(articleID uint) ([]entities.Comment, error) {
	var comments []entities.Comment
	result := r.db.Where("article_id = ?", articleID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (r *commentRepositoryImpl) CreateComment(comment entities.Comment) (entities.Comment, error) {
	result := r.db.Create(&comment)
	return comment, result.Error
}

func (r *commentRepositoryImpl) DeleteComment(id uint) error {
	result := r.db.Delete(&entities.Comment{}, id)
	return result.Error
}
