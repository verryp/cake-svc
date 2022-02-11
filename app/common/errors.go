package common

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
	ErrCreate = errors.New("failed created data")
)

func FormatValidationErrorResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
