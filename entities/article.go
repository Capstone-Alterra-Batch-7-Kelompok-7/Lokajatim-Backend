package entities

type Article struct {
    ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Title      string `json:"title"`
    Content    string `json:"content"`
    Photo      string `json:"photo"`
    Like       int    `json:"like"`
    Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
}
