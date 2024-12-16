package request

import "lokajatim/entities"

// CommentRequest is the request for the Comment endpoint
// @Description CommentRequest is the request for the Comment endpoint
// @Param UserID int true "ID of the user"
// @Param ArticleID int true "ID of the article"
// @Param Comment string true "Comment of the article"
type CommentRequest struct {
	UserID  	int    	`json:"user_id"`
	ArticleID 	int 	`json:"article_id"`
	Comment 	string 	`json:"comment"`
}

func (commentRequest CommentRequest) ToEntities() (entities.Comment, error) {
	return entities.Comment{
		UserID:  commentRequest.UserID,
		ArticleID: commentRequest.ArticleID,
		Comment: commentRequest.Comment,
	}, nil
}
