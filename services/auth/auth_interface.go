package services

import "lokajatim/entities"

type AuthInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetUserByID(userID int) (entities.User, error)
}
