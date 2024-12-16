package middleware

type JwtInterfaceReset interface {
	GenerateEmailJWT(email string) (string, error)
}
