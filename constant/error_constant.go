package constant

import "errors"

var (
	EMAIL_NOT_FOUND   = errors.New("email not found")
	EMAIL_IS_EMPTY    = errors.New("email kosong")
	PASSWORD_IS_EMPTY = errors.New("password kosong")
	INVALID_USER_ID	  = errors.New("invalid user id")
)
