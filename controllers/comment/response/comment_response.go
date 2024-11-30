package response

import (
	"lokajatim/entities"
	"time"
)

// CommentResponse is the response for the Comment controller
// @Description CommentResponse is the response for the Comment controller
// @Param ID int true "ID of the comment"
// @Param User User true "User of the comment"
// @Param Article Article true "Article of the comment"
// @Param Comment string true "Comment of the comment"
// @Param CreatedAt string true "Created At of the comment"
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
