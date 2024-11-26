package services

import (
	"errors"
	"lokajatim/constant"
	"lokajatim/entities"
	"lokajatim/middleware"
	"lokajatim/repositories/auth"

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
