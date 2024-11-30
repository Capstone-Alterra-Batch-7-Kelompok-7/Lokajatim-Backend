package response

import (
	"lokajatim/entities"
	"time"
)

type ArticleResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ArticleFromEntities(article entities.Article) ArticleResponse {
	return ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Photo:     article.Photo,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}
