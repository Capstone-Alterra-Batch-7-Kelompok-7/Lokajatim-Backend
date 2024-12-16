package comment

import "lokajatim/entities"

type CommentServiceInterface interface {
	GetCommentByID(id int) (entities.Comment, error)
	GetCommentsByArticleID(articleID int) ([]entities.Comment, error)
	CreateComment(comment entities.Comment) (entities.Comment, error)
	DeleteComment(id int) error
}