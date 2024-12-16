package auth

import "lokajatim/entities"

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

func FromEntities(user entities.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
