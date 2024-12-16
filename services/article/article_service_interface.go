package article

import "lokajatim/entities"

type ArticleService interface {
    GetAllArticles() ([]entities.Article, error)
    GetArticleByID(id int) (entities.Article, error)
    CreateArticle(article entities.Article) (entities.Article, error)
    UpdateArticle(id int, article entities.Article) (entities.Article, error)
    DeleteArticle(id int) error
}
