package entities

type Comment struct {
    ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleID uint   `json:"article_id"`
	Content   string `json:"content"`
}


