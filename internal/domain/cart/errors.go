package cart

import (
	"errors"
)

var (
	ErrItemAlreadyExistInCart = errors.New("Item already exist in cart")
	ErrCountInvalid           = errors.New("Count should be positive integer")
)
