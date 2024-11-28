package entities

type Comment struct {
    ID        	int   	`json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    	int   	`json:"user_id" gorm:"foreignKey:UserID"`
	User		User	`json:"User"`
	ArticleID   int   	`json:"article_id" gorm:"foreignKey:ArticleID"`
	Article		Article	`json:"Article"`
	Comment		string 	`json:"comment"`
}


