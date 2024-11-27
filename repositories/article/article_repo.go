package article

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type articleRepository struct{
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) GetAll() ([]entities.Article, error) {
	var articles []entities.Article
	
	result := r.db.Preload("Comments").Find(&articles) 
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (r *articleRepository) GetByID(id uint) (entities.Article, error) {
    var article entities.Article

    result := r.db.Preload("Comments").First(&article, id) 
    return article, result.Error
}

func (r *articleRepository) Create(article *entities.Article) (*entities.Article, error) {
	result := r.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	
	result = r.db.Preload("Comments").First(&article, article.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return article, nil
}

