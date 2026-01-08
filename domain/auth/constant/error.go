package constant

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrSessionNotFound = errors.New("session not found")

	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrWrongEmailOrPassword  = errors.New("wrong email or password")
	ErrInvalidAccessToken    = errors.New("invalid access token")
	ErrInvalidRefreshToken   = errors.New("invalid refresh token")
	ErrUserHasBeenBanned     = errors.New("user has been banned")
)
