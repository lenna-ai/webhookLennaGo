package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func StatusResponseErrorValidate(field string, tag string) string {
	return fmt.Sprintf("error field `%s` must be `%s`", field, tag)
}
