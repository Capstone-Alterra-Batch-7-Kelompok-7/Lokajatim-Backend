package auth_test

import (
	"errors"
	"testing"

	"lokajatim/entities"
	"lokajatim/services/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for AuthRepoInterface
type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) Login(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) GetUserByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) GetUserByID(userID int) (entities.User, error) {
	args := m.Called(userID)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) GetAllUsers() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockAuthRepo) Register(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) GetLastUserID() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockAuthRepo) UpdateUser(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockAuthRepo) DeleteUser(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockAuthRepo) StoreOTP(email, otp string) error {
	args := m.Called(email, otp)
	return args.Error(0)
}

func (m *MockAuthRepo) VerifyOTP(email, otp string) (bool, error) {
	args := m.Called(email, otp)
	return args.Bool(0), args.Error(1)
}

func (m *MockAuthRepo) UpdatePassword(user entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Mock for JwtInterface
type MockJwt struct {
	mock.Mock
}

func (m *MockJwt) GenerateJWT(userID int, name, role string) (string, error) {
	args := m.Called(userID, name, role)
	return args.String(0), args.Error(1)
}

func TestLogin(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	mockJwt := new(MockJwt)
	service := auth.NewAuthService(mockAuthRepo, mockJwt, nil)

	// Generate hashed password
	hashedPassword, _ := service.HashPassword("password")

	validUser := entities.User{
		Email:    "test@example.com",
		Password: "password",
	}
	storedUser := entities.User{
		ID:       1,
		Email:    "test@example.com",
		Password: hashedPassword,
		Name:     "Test User",
		Role:     "user",
	}

	mockAuthRepo.On("GetUserByEmail", validUser.Email).Return(storedUser, nil)
	mockJwt.On("GenerateJWT", storedUser.ID, storedUser.Name, storedUser.Role).Return("testtoken", nil)

	// Successful Login
	result, err := service.Login(validUser)
	assert.NoError(t, err)
	assert.Equal(t, "testtoken", result.Token)

	// Invalid Password
	mockAuthRepo.On("GetUserByEmail", validUser.Email).Return(storedUser, nil)
	invalidPasswordUser := entities.User{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}
	_, err = service.Login(invalidPasswordUser)
	assert.Error(t, err)
	assert.Equal(t, "incorrect password", err.Error())

	// Invalid Email
	mockAuthRepo.On("GetUserByEmail", "invalid@example.com").Return(entities.User{}, errors.New("user not found"))
	invalidUser := entities.User{
		Email:    "invalid@example.com",
		Password: "password",
	}
	_, err = service.Login(invalidUser)
	assert.Error(t, err)
}

func TestRegister(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	mockJwt := new(MockJwt)
	service := auth.NewAuthService(mockAuthRepo, mockJwt, nil)

	newUser := entities.User{
		Email:    "newuser@example.com",
		Password: "password",
	}

	mockAuthRepo.On("GetLastUserID").Return(1, nil)
	mockAuthRepo.On("Register", mock.Anything).Return(newUser, nil)
	mockJwt.On("GenerateJWT", mock.Anything, mock.Anything, mock.Anything).Return("testtoken", nil)

	// Successful Registration
	result, err := service.Register(newUser)
	assert.NoError(t, err)
	assert.Equal(t, "testtoken", result.Token)
}

func TestGetUserByID(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	service := auth.NewAuthService(mockAuthRepo, nil, nil)

	user := entities.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	mockAuthRepo.On("GetUserByID", user.ID).Return(user, nil)

	// Successful Get User By ID
	result, err := service.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	// User Not Found
	mockAuthRepo.On("GetUserByID", 2).Return(entities.User{}, errors.New("user not found"))
	_, err = service.GetUserByID(2)
	assert.Error(t, err)
}

func TestGetAllUsers(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	service := auth.NewAuthService(mockAuthRepo, nil, nil)

	users := []entities.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
	}

	mockAuthRepo.On("GetAllUsers").Return(users, nil)

	// Successful Get All Users
	result, err := service.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, users, result)
}

func TestUpdateUser(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	service := auth.NewAuthService(mockAuthRepo, nil, nil)

	existingUser := entities.User{
		ID:    1,
		Name:  "Old Name",
		Email: "test@example.com",
	}

	updatedData := entities.User{
		Name: "New Name",
	}

	updatedUser := entities.User{
		ID:    1,
		Name:  "New Name",
		Email: "test@example.com",
	}

	mockAuthRepo.On("GetUserByID", existingUser.ID).Return(existingUser, nil)
	mockAuthRepo.On("UpdateUser", mock.Anything).Return(updatedUser, nil)

	// Successful Update User
	result, err := service.UpdateUser(existingUser.ID, updatedData)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, result)
}

func TestDeleteUser(t *testing.T) {
	mockAuthRepo := new(MockAuthRepo)
	service := auth.NewAuthService(mockAuthRepo, nil, nil)

	mockAuthRepo.On("DeleteUser", 1).Return(nil)

	// Successful Delete User
	err := service.DeleteUser(1)
	assert.NoError(t, err)
}
