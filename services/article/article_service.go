package article

import (
	"lokajatim/entities"
	"lokajatim/repositories/article"
)

type ArticleService struct {
	articleRepository article.ArticleRepository
}

func NewArticleService(articleRepo article.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository: articleRepo}
}

func (s *ArticleService) GetAllArticles() ([]entities.Article, error) {
	return s.articleRepository.GetAll()
}

func (s *ArticleService) GetArticleByID(id int) (entities.Article, error) {
	return s.articleRepository.GetByID(id)
}

func (s *ArticleService) CreateArticle(article entities.Article) (entities.Article, error) {
	return s.articleRepository.Create(article)
}

func (s *ArticleService) UpdateArticle(id int, article entities.Article) (entities.Article, error) {
	return s.articleRepository.Update(id, article)
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.articleRepository.Delete(id)
}
