package response

import "lokajatim/entities"

type CommentResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Comment   string `json:"comment"`
}

func CommentFromEntities(comment entities.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		Comment:   comment.Comment,
		}
}