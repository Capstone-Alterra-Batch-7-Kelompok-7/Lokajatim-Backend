package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari konteks
		user := c.Get("user").(*jwt.Token)

		// Ambil klaim khusus dari token
		claims, ok := user.Claims.(*JwtCustomClaims)
		if !ok {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Access forbidden: invalid claims type",
			})
		}

		// Periksa apakah peran adalah "admin"
		if claims.Role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Access forbidden: admin only",
			})
		}

		return next(c)
	}
}
