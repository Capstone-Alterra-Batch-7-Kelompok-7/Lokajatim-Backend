package middleware

type JwtInterface interface {
	GenerateJWT(userID int, name, role string) (string, error)
}
