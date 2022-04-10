package user

import (
	"errors"
)

var (
	ErrUserExistWithName = errors.New("Username already exist")
	ErrUserNotFound      = errors.New("User not found")
	//TODO: Validation error can be detailed
	ErrMismatchedPasswords = errors.New("Given passwords does not match")
	ErrInvalidUsername     = errors.New("Invalid username")
	ErrInvalidPassword     = errors.New("Invalid password")
)
