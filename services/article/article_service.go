package article

import (
	"lokajatim/entities"
	"lokajatim/repositories/article"
)

type ArticleService interface {
    GetAllArticles() ([]entities.Article, error)
    GetArticleByID(id uint) (entities.Article, error)
    CreateArticle(article *entities.Article) (*entities.Article, error)
}

type articleService struct {
    articleRepo article.ArticleRepository
}

func NewArticleService(repo article.ArticleRepository) ArticleService {
    return &articleService{articleRepo: repo}
}

func (s *articleService) GetAllArticles() ([]entities.Article, error) {
    return s.articleRepo.GetAll()
}

func (s *articleService) GetArticleByID(id uint) (entities.Article, error) {
    return s.articleRepo.GetByID(id)
}

func (s *articleService) CreateArticle(article *entities.Article) (*entities.Article, error) {
	return s.articleRepo.Create(article)
}
