package article

import (
	"lokajatim/entities"
)

type ArticleRepository interface {
	GetAll() ([]entities.Article, error)
	GetByID(id int) (entities.Article, error)
	Create(article entities.Article) (entities.Article, error)
	Update(id int, article entities.Article) (entities.Article, error)
	Delete(id int) error
}
