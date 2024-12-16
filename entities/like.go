package entities

import "time"

type Like struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleID int       `json:"article_id" gorm:"foreignKey:ArticleID"`
	Article   Article   `json:"Article"`
	UserID    int       `json:"user_id" gorm:"foreignKey:UserID"`
	User      User      `json:"User"`
	CreatedAt time.Time `json:"created_at"`
}