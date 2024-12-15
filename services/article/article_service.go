package article

import (
	"lokajatim/entities"
	"lokajatim/repositories/article"
)

type ArticleServiceImpl struct {
	articleRepository article.ArticleRepository
}

func NewArticleService(articleRepo article.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{articleRepository: articleRepo}
}

func (s *ArticleServiceImpl) GetAllArticles() ([]entities.Article, error) {
	return s.articleRepository.GetAll()
}

func (s *ArticleServiceImpl) GetArticleByID(id int) (entities.Article, error) {
	return s.articleRepository.GetByID(id)
}

func (s *ArticleServiceImpl) CreateArticle(article entities.Article) (entities.Article, error) {
	return s.articleRepository.Create(article)
}

func (s *ArticleServiceImpl) UpdateArticle(id int, article entities.Article) (entities.Article, error) {
	return s.articleRepository.Update(id, article)
}

func (s *ArticleServiceImpl) DeleteArticle(id int) error {
	return s.articleRepository.Delete(id)
}
