package user

import (
	"errors"
)

var (
	ErrUserExistWithName = errors.New("User already exist with same username in database")
	ErrUserNotFound      = errors.New("User not found")
	//TODO: Validation error can be specified
	ErrMismatchedPasswords = errors.New("Given passwords does not match")
	ErrInvalidUsername     = errors.New("Invalid username")
	ErrInvalidPassword     = errors.New("Invalid password")
)
