package comment

import "lokajatim/entities"

type CommentRepository interface {
	GetAllComments() ([]entities.Comment, error)
	CreateComment(comment entities.Comment) (entities.Comment, error)
	DeleteComment(id uint) error
}
