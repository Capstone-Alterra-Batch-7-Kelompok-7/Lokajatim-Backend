package request

import "lokajatim/entities"


// ArticleRequest is the request for the Article endpoint
// @Description ArticleRequest is the request for the Article endpoint
// @Param Title string true "Title of the article"
// @Param Content string true "Content of the article"
// @Param Photo string true "Photo of the article"
type ArticleRequest struct {
	Title  string `json:"title"`
	Content string `json:"content"`
	Photo string `json:"photo"`
}

func (articleRequest ArticleRequest) ToEntities() (entities.Article, error) {
	return entities.Article {
		Title: articleRequest.Title,
		Content: articleRequest.Content,
		Photo: articleRequest.Photo,
	}, nil
}