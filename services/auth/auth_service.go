package auth

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

// AuthService struct untuk mengimplementasikan layanan autentikasi
type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
	jwtInterface      middleware.JwtInterface
}

// NewAuthService untuk membuat instance baru AuthService
func NewAuthService(ar auth.AuthRepoInterface, jt middleware.JwtInterface, jtr middleware.JwtInterfaceReset) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

// Login memverifikasi email dan password, lalu mengembalikan user dengan token
func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	}
	if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	dbUser, err := authService.authRepoInterface.GetUserByEmail(user.Email)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	if !authService.CheckPasswordHash(user.Password, dbUser.Password) {
		return entities.User{}, errors.New("incorrect password")
	}

	token, err := authService.jwtInterface.GenerateJWT(dbUser.ID, dbUser.Name, dbUser.Role)
	if err != nil {
		return entities.User{}, errors.New("failed to generate token")
	}
	dbUser.Token = token

	return dbUser, nil
}

// Register membuat user baru dan menghasilkan token
func (authService AuthService) Register(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	}
	if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	user.Role = "user" // Default role

	hashedPassword, err := authService.HashPassword(user.Password)
	if err != nil {
		return entities.User{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	lastID, err := authService.authRepoInterface.GetLastUserID()
	if err != nil {
		return entities.User{}, errors.New("failed to fetch last user ID")
	}
	user.ID = lastID + 1

	createdUser, err := authService.authRepoInterface.Register(user)
	if err != nil {
		return entities.User{}, errors.New("failed to register user")
	}

	token, err := authService.jwtInterface.GenerateJWT(createdUser.ID, createdUser.Name, createdUser.Role)
	if err != nil {
		return entities.User{}, errors.New("failed to generate token")
	}
	createdUser.Token = token

	return createdUser, nil
}

// GetUserByID mengambil user berdasarkan ID
func (authService AuthService) GetUserByID(userID int) (entities.User, error) {
	if userID <= 0 {
		return entities.User{}, errors.New("invalid user ID")
	}

	user, err := authService.authRepoInterface.GetUserByID(userID)
	if err != nil {
		return entities.User{}, errors.New("user not found")
	}

	return user, nil
}

// GetAllUsers mengambil semua data user
func (authService AuthService) GetAllUsers() ([]entities.User, error) {
	users, err := authService.authRepoInterface.GetAllUsers()
	if err != nil {
		return nil, errors.New("failed to fetch users")
	}

	return users, nil
}

// UpdateUser memperbarui data user berdasarkan ID
func (authService AuthService) UpdateUser(userID int, updatedData entities.User) (entities.User, error) {
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

	updatedUser, err := authService.authRepoInterface.UpdateUser(existingUser)
	if err != nil {
		return entities.User{}, errors.New("failed to update user")
	}

	return updatedUser, nil
}

func (authService AuthService) DeleteUser(userID int) error {
	err := authService.authRepoInterface.DeleteUser(userID)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (authService *AuthService) SendOTPToEmail(email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	_, err := authService.authRepoInterface.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("email not found")
	}

	otp := utils.GenerateOTP()

	if err := authService.authRepoInterface.StoreOTP(email, otp); err != nil {
		return "", errors.New("failed to store OTP")
	}

	if err := utils.SendOTPEmail(email, otp); err != nil {
		return "", errors.New("failed to send OTP email")
	}

	return "OTP successfully sent", nil
}

func (authService *AuthService) ResetPassword(email, otp, newPassword string) (string, error) {
	valid, err := authService.authRepoInterface.VerifyOTP(email, otp)
	if err != nil || !valid {
		return "", errors.New("invalid OTP")
	}

	hashedPassword, err := authService.HashPassword(newPassword)
	if err != nil {
		return "", errors.New("failed to hash new password")
	}

	user, _ := authService.authRepoInterface.GetUserByEmail(email)
	user.Password = hashedPassword

	if err := authService.authRepoInterface.UpdatePassword(user); err != nil {
		return "", errors.New("failed to update password")
	}

	return "Password successfully updated", nil
}

// HashPassword hashes a password
func (authService *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash checks a hashed password
func (authService *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (authService *AuthService) VerifyOTP(email, otp string) (bool, error) {
	// Get the OTP and expiration time from the database
	user, err := authService.authRepoInterface.GetUserByEmail(email)
    if err != nil {
        return false, fmt.Errorf("user not found")
    }

    // Cek kesesuaian OTP
    if user.OTP != otp {
        return false, fmt.Errorf("invalid OTP, ini servis")
    }

    return true, nil
}

func (s *AuthService) GetUserByEmail(email string) (*entities.User, error) {
    user, err := s.authRepoInterface.GetUserByEmail(email)
    if err != nil {
        return nil, err 
    }
    return &user, nil
}

func (s *AuthService) StoreOTP(email, otp string) error {
    if email == "" || otp == "" {
        return errors.New("email and otp cannot be empty")
    }

    err := s.authRepoInterface.StoreOTP(email, otp)
    if err != nil {
        return errors.New("failed to store OTP in database")
    }

    return nil
}