package request

import "lokajatim/entities"

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