package comment

import (
	"lokajatim/entities"
	"lokajatim/repositories/comment"
)

type CommentService struct {
	commentRepository comment.CommentRepository
}

func NewCommentService(commentRepo comment.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepo}
}

func (s *CommentService) GetCommentByID(id uint) (entities.Comment, error) {
	return s.commentRepository.GetCommentByID(id)
}

func (s *CommentService) GetCommentsByArticleID(articleID uint) ([]entities.Comment, error) {
	return s.commentRepository.GetCommentsByArticleID(articleID)
}

func (s *CommentService) CreateComment(comment entities.Comment) (entities.Comment, error) {
	createdComment, err := s.commentRepository.CreateComment(comment)
	if err != nil {
		return entities.Comment{}, err
	}
	return createdComment, nil
}

func (s *CommentService) DeleteComment(id uint) error {
	return s.commentRepository.DeleteComment(id)
}
