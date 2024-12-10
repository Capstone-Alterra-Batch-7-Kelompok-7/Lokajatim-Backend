package auth

import (
	"errors"
	"lokajatim/controllers/auth/request"
	"lokajatim/controllers/auth/response"
	"lokajatim/controllers/base"
	services "lokajatim/services/auth"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as services.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService services.AuthInterface
}

// @Summary Login
// @Description User login endpoint
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} base.BaseResponse
// @Router /login [post]
func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	if err := c.Bind(&userLogin); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}

	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"email": userLogin.Email,
		})
	}

	return base.SuccesResponse(c, response.LoginFromEntities(user))
}

// @Summary Register
// @Description User registration endpoint
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Register Request"
// @Success 201 {object} response.RegisterResponse
// @Failure 400 {object} base.BaseResponse
// @Router /register [post]
func (userController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	if err := c.Bind(&userRegister); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request payload",
		})
	}

	user, err := userController.authService.Register(userRegister.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"email": userRegister.Email,
		})
	}

	return base.SuccesResponse(c, response.RegisterFromEntities(user))
}

// GetUserByID handles the request to retrieve a user by their ID.
// @Summary Get User by ID
// @Description Retrieve details of a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} base.BaseResponse{data=response.UserResponse}
// @Failure 400 {object} base.BaseResponse
// @Failure 404 {object} base.BaseResponse
// @Router /users/{id} [get]
func (userController AuthController) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil || userID <= 0 {
		return base.ErrorResponse(c, errors.New("invalid user ID"), nil)
	}

	user, err := userController.authService.GetUserByID(userID)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	response := response.UpdateFromEntities(user)
	return base.SuccesResponse(c, response)
}

// SendOTPController handles the send OTP request
// @Summary Send OTP to the user's email
// @Description This endpoint sends an OTP to the provided email address
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.SendOTPRequest true "Request to send OTP"
// @Success 200 {object} response.SendOTPResponse "OTP sent successfully"
// @Failure 400 {object} base.BaseResponse "Invalid email address"
// @Failure 500 {object} base.BaseResponse "Internal server error"
// @Router /forgot-password [post]
func (userController *AuthController) SendOTPController(c echo.Context) error {
	var request request.SendOTPRequest
	if err := c.Bind(&request); err != nil {
		return base.ErrorResponse(c, err, "Failed to bind request parameters")
	}

	// Call service to send OTP
	message, err := userController.authService.SendOTPToEmail(request.Email)
	if err != nil {
		return base.ErrorResponse(c, err, "Failed to send OTP to email adress")
	}

	// Return success response
	response := response.SendOTPResponse{Message: message}
	return base.SuccesResponse(c, response)
}

// ResetPasswordController handles the password reset request
// @Summary Reset the user's password
// @Description This endpoint resets the password after validating the OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.ResetPasswordRequest true "Request to reset password"
// @Success 200 {object} response.ResetPasswordResponse "Password reset successfully"
// @Failure 400 {object} base.BaseResponse "Invalid OTP or email"
// @Failure 500 {object} base.BaseResponse "Internal server error"
// @Router /reset-password [post]
func (userController *AuthController) ResetPasswordController(c echo.Context) error {
	var request request.ResetPasswordRequest
	if err := c.Bind(&request); err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	// Call service to reset password
	message, err := userController.authService.ResetPassword(request.Email, request.OTP, request.NewPassword)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	// Return success response
	response := response.ResetPasswordResponse{Message: message}
	return base.SuccesResponse(c, response)
}

// GetAllUsersController handles fetching all users
// @Summary Get All Users
// @Description Fetch a list of all users
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {array} response.RegisterResponse "List of users"
// @Failure 500 {object} base.BaseResponse "Internal server error"
// @Router /users [get]
func (userController *AuthController) GetAllUsersController(c echo.Context) error {
	users, err := userController.authService.GetAllUsers()
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	var userResponses []response.UpdateUserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.UpdateFromEntities(user))
	}

	return base.SuccesResponse(c, userResponses)
}

// UpdateUserController handles updating a user's data
// @Summary Update User
// @Description Update a user's information by ID
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body request.UpdateUserRequest true "Update User Request"
// @Success 200 {object} response.RegisterResponse "User updated successfully"
// @Failure 400 {object} base.BaseResponse "Invalid input or user not found"
// @Failure 500 {object} base.BaseResponse "Internal server error"
// @Router /users/{id} [put]
func (userController *AuthController) UpdateUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil || userID <= 0 {
		return base.ErrorResponse(c, errors.New("invalid user ID"), nil)
	}

	var updateRequest request.UpdateUserRequest
	if err := c.Bind(&updateRequest); err != nil {
		return base.ErrorResponse(c, err, "Failed to bind request parameters")
	}

	updatedUser, err := userController.authService.UpdateUser(userID, updateRequest.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	response := response.UpdateFromEntities(updatedUser)
	return base.SuccesResponse(c, response)
}

// DeleteUserController handles deleting a user
// @Summary Delete User
// @Description Delete a user by ID
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} base.BaseResponse "User deleted successfully"
// @Failure 400 {object} base.BaseResponse "Invalid user ID"
// @Failure 500 {object} base.BaseResponse "Internal server error"
// @Router /users/{id} [delete]
func (ac *AuthController) DeleteUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil || userID <= 0 {
		return base.ErrorResponse(c, errors.New("invalid user ID"), nil)
	}

	err = ac.authService.DeleteUser(userID)
	if err != nil {
		return base.ErrorResponse(c, err, nil)
	}

	return base.SuccesResponse(c, map[string]string{"message": "User deleted successfully"})
}

