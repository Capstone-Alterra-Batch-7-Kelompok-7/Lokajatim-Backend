package article

import (
	"lokajatim/entities"
)

type ArticleRepository interface {
	GetAll() ([]entities.Article, error)
	GetByID(id uint) (entities.Article, error)
	Create(article *entities.Article) (*entities.Article, error)
	Update(id uint, article *entities.Article) (*entities.Article, error)
	Delete(id uint) error
}
