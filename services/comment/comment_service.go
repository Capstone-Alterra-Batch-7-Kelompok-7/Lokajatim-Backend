package comment

import (
	"lokajatim/entities"
	"lokajatim/repositories/comment"
)

type CommentService struct {
	commentRepository comment.CommentRepositoryInterface
}

func NewCommentService(commentRepo comment.CommentRepositoryInterface) *CommentService {
	return &CommentService{commentRepository: commentRepo}
}

func (s *CommentService) GetCommentByID(id int) (entities.Comment, error) {
	return s.commentRepository.GetCommentByID(id)
}

func (s *CommentService) GetCommentsByArticleID(articleID int) ([]entities.Comment, error) {
	return s.commentRepository.GetCommentsByArticleID(articleID)
}

func (s *CommentService) CreateComment(comment entities.Comment) (entities.Comment, error) {
	return s.commentRepository.CreateComment(comment)
}

func (s *CommentService) DeleteComment(id int) error {
	return s.commentRepository.DeleteComment(id)
}
