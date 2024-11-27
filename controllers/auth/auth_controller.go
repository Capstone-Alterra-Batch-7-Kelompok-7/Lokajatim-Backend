package auth

import (
	"lokajatim/controllers/auth/request"
	"lokajatim/controllers/auth/response"
	"lokajatim/controllers/base"
	services "lokajatim/services/auth"

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
