package article

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type articleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepositoryImpl{db: db}
}

func (r *articleRepositoryImpl) GetAll() ([]entities.Article, error) {
	var articles []entities.Article
	if err := r.db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepositoryImpl) GetByID(id int) (entities.Article, error) {
	var article entities.Article
	result := r.db.First(&article, id)
	if result.Error != nil {
		return entities.Article{}, result.Error
	}
	return article, nil
}

func (r *articleRepositoryImpl) Create(article entities.Article) (entities.Article, error) {
	if err := r.db.Create(&article).Error; err != nil {
		return entities.Article{}, err
	}
	return article, nil
}

func (r *articleRepositoryImpl) Update(id int, article entities.Article) (entities.Article, error) {
	if err := r.db.Model(&entities.Article{}).Where("id = ?", id).Updates(article).Error; err != nil {
        return entities.Article{}, err
    }

    var updatedArticle entities.Article
    if err := r.db.First(&updatedArticle, id).Error; err != nil {
        return entities.Article{}, err
    }

    return updatedArticle, nil
}

func (r *articleRepositoryImpl) Delete(id int) error {
	var article entities.Article
	if err := r.db.Delete(&article, id).Error; err != nil {
		return err
	}
	return nil
}
