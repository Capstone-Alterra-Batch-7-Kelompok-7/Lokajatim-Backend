package entities

type Article struct {
    ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
    Title       string  `json:"title"`
    Content     string  `json:"content"`
    Photo       string  `json:"photo"`
}
