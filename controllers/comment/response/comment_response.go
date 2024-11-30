package response

import (
	"lokajatim/entities"
	"time"
)

type CommentResponse struct {
	ID        int              `json:"id"`
	User      entities.User    `json:"user"`
	Article   entities.Article `json:"article"`
	Comment   string           `json:"comment"`
	CreatedAt time.Time        `json:"created_at"`
}

func CommentFromEntities(comment entities.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		User:      comment.User,
		Article:   comment.Article,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
	}
}
