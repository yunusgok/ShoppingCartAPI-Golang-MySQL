package product

import (
	"errors"
)

var (
	ErrProductNotFound         = errors.New("Product not found")
	ErrProductStockIsNotEnough = errors.New("Product stock is not enough")
)
