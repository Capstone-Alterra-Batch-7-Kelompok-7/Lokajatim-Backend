package response

import "lokajatim/entities"

type ArticleResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
}

func ArticleFromEntities(article entities.Article) ArticleResponse {
	return ArticleResponse {
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Photo:   article.Photo,
	}
}