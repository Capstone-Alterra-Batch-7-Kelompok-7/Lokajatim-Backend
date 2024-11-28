package request

import "lokajatim/entities"

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
