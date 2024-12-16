package auth

import "lokajatim/entities"

type AuthServiceInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetUserByID(userID int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	DeleteUser(userID int) error
	UpdateUser(userID int, updatedData entities.User) (entities.User, error)
	SendOTPToEmail(email string) (string, error)
	ResetPassword(email, otp, newPassword string) (string, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	VerifyOTP(email, otp string) (bool, error) 
	GetUserByEmail(email string) (*entities.User, error)
	StoreOTP(email, otp string) error
}
