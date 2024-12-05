package auth

import "lokajatim/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetLastUserID() (int, error)
	GetUserByID(userID int) (entities.User, error)

	StoreOTP(email string, otp string) error
	VerifyOTP(email string, otp string) (bool, error)
	UpdatePassword(user entities.User) error
}
