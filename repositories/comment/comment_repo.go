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

func (r *commentRepositoryImpl) GetAllComments() ([]entities.Comment, error) {
	var comments []entities.Comment
	result := r.db.Find(&comments)
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
