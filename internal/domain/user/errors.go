package user

import (
	"errors"
)

var (
	ErrUserExistWithName = errors.New("User already exist with same username in database")
	ErrUserNotFound      = errors.New("User not found")
)
