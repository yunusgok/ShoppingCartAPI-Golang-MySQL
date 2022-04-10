package api_helper

import "errors"

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

var (
	ErrInvalidBody = errors.New("Check your request body")
)
