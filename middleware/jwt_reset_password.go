package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JwtCustomClaimsReset now only includes the email in the payload
type JwtCustomClaimsReset struct {
	Email string `json:"email"` // Only email is included in the JWT payload
	jwt.RegisteredClaims
}

type JwtLokajatimReset struct {
}

// GenerateJWT for reset password now only requires the email
func (jwtLokajatim JwtLokajatimReset) GenerateEmailJWT(email string) (string, error) {
	// Only email is passed in the claims
	claims := &JwtCustomClaimsReset{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 3 days expiration
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func JWTMiddlewareReset(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		token = strings.TrimPrefix(token, "Bearer ")

		parsedToken, err := jwt.ParseWithClaims(token, &JwtCustomClaimsReset{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !parsedToken.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := parsedToken.Claims.(*JwtCustomClaimsReset); ok {
			if claims.ExpiresAt.Time.Before(time.Now()) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired")
			}

			c.Set("email", claims.Email)
		}

		return next(c)
	}
}
