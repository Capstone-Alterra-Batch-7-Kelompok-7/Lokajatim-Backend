package response

import (
	"lokajatim/entities"
	"time"
)

// ArticleResponse is the response for the Article controller
// @Description ArticleResponse is the response for the Article controller
// @Param ID int true "ID of the article"
// @Param Title string true "Title of the article"
// @Param Content string true "Content of the article"
// @Param Photo string true "Photo of the article"
// @Param CreatedAt string true "Created At of the article"
// @Param UpdatedAt string true "Updated At of the article"
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
