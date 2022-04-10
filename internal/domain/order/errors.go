package order

import "errors"

var (
	ErrEmptyCartFound       = errors.New("Cart is empty")
	ErrInvalidOrderID       = errors.New("Invalid Order ID")
	ErrCancelDurationPassed = errors.New("Cancel duration passed")
	ErrNotEnoughStock       = errors.New("Not enough stock")
)
