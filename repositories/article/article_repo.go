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

func (r *articleRepository) Update(id uint, article *entities.Article) (*entities.Article, error) {
	var existingArticle entities.Article
	result := r.db.First(&existingArticle, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update article fields
	existingArticle.Title = article.Title
	existingArticle.Content = article.Content
	existingArticle.Photo = article.Photo
	existingArticle.Like = article.Like

	// Save the updated article
	result = r.db.Save(&existingArticle)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingArticle, nil
}

func (r *articleRepository) Delete(id uint) error {
	result := r.db.Delete(&entities.Article{}, id)
	return result.Error
}
