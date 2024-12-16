package base

import (
	"lokajatim/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`      // Response data for success
	ErrorData interface{} `json:"error_data,omitempty"` // Detailed error data for failure
}

// SuccesResponse generates a successful response
func SuccesResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

// ErrorResponse generates an error response with optional error data
func ErrorResponse(c echo.Context, err error, errorData any) error {
	return c.JSON(helper.GetResponseCodeFromErr(err), BaseResponse{
		Status:    false,
		Message:   err.Error(),
		ErrorData: errorData,
	})
}
