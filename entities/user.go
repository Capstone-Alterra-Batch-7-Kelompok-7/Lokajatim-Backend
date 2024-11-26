package entities

type User struct {
	ID          int    `gorm:"primaryKey" json:"id_user"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
    Password    string `json:"password"`
	Address     string `json:"address" gorm:"null"`
	PhoneNumber string `json:"phone_number" goem:"null"`
	NIK         string `json:"nik"`
	Token       string `json:"token"`
	Role        string `json:"role"`
}
