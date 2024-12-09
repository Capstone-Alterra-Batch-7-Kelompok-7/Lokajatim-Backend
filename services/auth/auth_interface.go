package services

import "lokajatim/entities"

type AuthInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetUserByID(userID int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(userID int, updatedData entities.User) (entities.User, error)
	DeleteUser(userID int) error

	SendOTPToEmail(email string) (string, error)
	ResetPassword(email string, otp string, newPassword string) (string, error)
}
