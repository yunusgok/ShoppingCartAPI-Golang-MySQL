package category

import (
	"errors"
)

var (
	ErrCategoryExistWithName = errors.New("Category already exist with same name in database")
)
