package auth_test

import (
	// "testing"

	"lokajatim/entities"
	// "lokajatim/services/auth"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthRepo untuk AuthRepoInterface
type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) GetUserByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	user, _ := args.Get(0).(entities.User)
	return user, args.Error(1)
}

func (m *MockAuthRepo) GetLastUserID() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockAuthRepo) Register(user entities.User) (entities.User, error) {
	args := m.Called(user)
	registeredUser, _ := args.Get(0).(entities.User)
	return registeredUser, args.Error(1)
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

func (m *MockAuthRepo) GetUserByID(userID int) (entities.User, error) {
	args := m.Called(userID)
	user, _ := args.Get(0).(entities.User)
	return user, args.Error(1)
}

func (m *MockAuthRepo) GetAllUsers() ([]entities.User, error) {
	args := m.Called()
	users, _ := args.Get(0).([]entities.User)
	return users, args.Error(1)
}

func (m *MockAuthRepo) UpdateUser(user entities.User) (entities.User, error) {
	args := m.Called(user)
	updatedUser, _ := args.Get(0).(entities.User)
	return updatedUser, args.Error(1)
}

func (m *MockAuthRepo) DeleteUser(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockAuthRepo) Login(user entities.User) (entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(entities.User), args.Error(1)
}
// MockJwt untuk JwtInterface
type MockJwt struct {
	mock.Mock
}

func (m *MockJwt) GenerateJWT(userID int, name, role string) (string, error) {
	args := m.Called(userID, name, role)
	return args.String(0), args.Error(1)
}

// Test Login sukses
// func TestLogin_Success(t *testing.T) {
// 	mockRepo := new(MockAuthRepo)
// 	mockJwt := new(MockJwt)
// 	authService := auth.NewAuthService(mockRepo, mockJwt)

// 	userInput := entities.User{Email: "user@example.com", Password: "password123"}
// 	dbUser := entities.User{ID: 1, Email: "user@example.com", Password: "$2a$14$hashedPassword", Name: "Test User", Role: "user"}

// 	mockRepo.On("GetUserByEmail", userInput.Email).Return(dbUser, nil)
// 	mockJwt.On("GenerateJWT", dbUser.ID, dbUser.Name, dbUser.Role).Return("valid_jwt_token", nil)

// 	result, err := authService.Login(userInput)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "valid_jwt_token", result.Token)
// 	mockRepo.AssertExpectations(t)
// 	mockJwt.AssertExpectations(t)
// }

// // Test Register sukses
// func TestRegister_Success(t *testing.T) {
// 	mockRepo := new(MockAuthRepo)
// 	mockJwt := new(MockJwt)
// 	authService := auth.NewAuthService(mockRepo, mockJwt)

// 	userInput := entities.User{Email: "newuser@example.com", Password: "password123", Name: "New User"}
// 	registeredUser := entities.User{ID: 1, Email: "newuser@example.com", Password: "$2a$14$hashedPassword", Name: "New User"}

// 	mockRepo.On("GetLastUserID").Return(0, nil)
// 	mockRepo.On("Register", mock.Anything).Return(registeredUser, nil)
// 	mockJwt.On("GenerateJWT", registeredUser.ID, registeredUser.Name, registeredUser.Role).Return("valid_jwt_token", nil)

// 	result, err := authService.Register(userInput)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "valid_jwt_token", result.Token)
// 	mockRepo.AssertExpectations(t)
// 	mockJwt.AssertExpectations(t)
// }

// Test GetUserByID sukses
// func TestGetUserByID_Success(t *testing.T) {
// 	mockRepo := new(MockAuthRepo)
// 	authService := auth.NewAuthService(mockRepo, nil)

// 	dbUser := entities.User{ID: 1, Email: "user@example.com", Name: "Test User"}

// 	mockRepo.On("GetUserByID", 1).Return(dbUser, nil)

// 	result, err := authService.GetUserByID(1)

// 	assert.NoError(t, err)
// 	assert.Equal(t, dbUser, result)
// 	mockRepo.AssertExpectations(t)
// }

// Test GetAllUsers sukses
// func TestGetAllUsers_Success(t *testing.T) {
// 	mockRepo := new(MockAuthRepo)
// 	authService := auth.NewAuthService(mockRepo, nil)

// 	users := []entities.User{
// 		{ID: 1, Email: "user1@example.com"},
// 		{ID: 2, Email: "user2@example.com"},
// 	}

// 	mockRepo.On("GetAllUsers").Return(users, nil)

// 	result, err := authService.GetAllUsers()

// 	assert.NoError(t, err)
// 	assert.Equal(t, users, result)
// 	mockRepo.AssertExpectations(t)
// }

// Test DeleteUser sukses
// func TestDeleteUser_Success(t *testing.T) {
// 	mockRepo := new(MockAuthRepo)
// 	authService := auth.NewAuthService(mockRepo, nil)

// 	mockRepo.On("DeleteUser", 1).Return(nil)

// 	err := authService.DeleteUser(1)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }
