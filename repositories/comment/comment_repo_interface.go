package comment

import "lokajatim/entities"

type CommentRepository interface {
	GetCommentByID(id uint) (entities.Comment, error)
	GetCommentsByArticleID(articleID uint) ([]entities.Comment, error)
	CreateComment(comment entities.Comment) (entities.Comment, error)
	DeleteComment(id uint) error
}
