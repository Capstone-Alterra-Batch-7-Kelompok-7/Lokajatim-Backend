package services

import (
	"errors"
	"fmt"
	"lokajatim/constant"
	"lokajatim/entities"
	"lokajatim/middleware"
	"lokajatim/repositories/auth"
	"lokajatim/utils"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar auth.AuthRepoInterface, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	// Retrieve user from database
	dbUser, err := authService.authRepoInterface.GetUserByEmail(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	// Check if the provided password matches the hashed password
	if !CheckPasswordHash(user.Password, dbUser.Password) {
		return entities.User{}, errors.New("incorrect password")
	}

	// Generate JWT token for the user
	token, err := authService.jwtInterface.GenerateJWT(dbUser.ID, dbUser.Name, dbUser.Role)
	if err != nil {
		return entities.User{}, err
	}
	dbUser.Token = token

	return dbUser, nil
}

func (authService AuthService) Register(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Hash password before saving to database
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = hashedPassword

	// Get the last ID and assign a new ID for the user
	lastID, err := authService.authRepoInterface.GetLastUserID()
	if err != nil {
		return entities.User{}, err
	}
	user.ID = lastID + 1

	// Register new user in the database
	createdUser, err := authService.authRepoInterface.Register(user)
	if err != nil {
		return entities.User{}, err
	}

	// Generate JWT token for the new user
	token, err := authService.jwtInterface.GenerateJWT(createdUser.ID, createdUser.Name, createdUser.Role)
	if err != nil {
		return entities.User{}, err
	}
	createdUser.Token = token

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (authService AuthService) GetUserByID(userID int) (entities.User, error) {
	if userID <= 0 {
		return entities.User{}, errors.New("invalid user ID")
	}

	// Ambil user dari database menggunakan authRepoInterface
	user, err := authService.authRepoInterface.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (authService *AuthService) SendOTPToEmail(email string) (string, error) {
	// Cek apakah email ada
	_, err := authService.authRepoInterface.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("email tidak ditemukan")
	}

	// Generate OTP
	otp := utils.GenerateOTP()

	// Simpan OTP ke database
	if err := authService.authRepoInterface.StoreOTP(email, otp); err != nil {
		return "", err
	}

	// Kirim OTP via email
	if err := utils.SendOTPEmail(email, otp); err != nil {
		return "", err
	}

	return "OTP berhasil dikirim", nil
}

func (authService *AuthService) ResetPassword(email string, otp string, newPassword string) (string, error) {
	// Verifikasi OTP
	valid, err := authService.authRepoInterface.VerifyOTP(email, otp)
	if err != nil || !valid {
		return "", fmt.Errorf("OTP tidak valid")
	}

	// Hash password baru
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	// Update password user
	user, _ := authService.authRepoInterface.GetUserByEmail(email)
	user.Password = hashedPassword

	if err := authService.authRepoInterface.UpdatePassword(user); err != nil {
		return "", err
	}

	return "Password berhasil diperbarui", nil
}

func (authService *AuthService) GetAllUsers() ([]entities.User, error) {
	return authService.authRepoInterface.GetAllUsers()
}

func (authService *AuthService) UpdateUser(userID int, updatedData entities.User) (entities.User, error) {
	existingUser, err := authService.authRepoInterface.GetUserByID(userID)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	if updatedData.Name != "" {
		existingUser.Name = updatedData.Name
	}
	if updatedData.Address != "" {
		existingUser.Address = updatedData.Address
	}
	if updatedData.PhoneNumber != "" {
		existingUser.PhoneNumber = updatedData.PhoneNumber
	}
	if updatedData.NIK != "" {
		existingUser.NIK = updatedData.NIK
	}

	// Save the updated user in the database
	updatedUser, err := authService.authRepoInterface.UpdateUser(existingUser)
	if err != nil {
		return entities.User{}, errors.New("failed to update user")
	}

	return updatedUser, nil
}

func (authService *AuthService) DeleteUser(userID int) error {
	return authService.authRepoInterface.DeleteUser(userID)
}