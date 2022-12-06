package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"failed_field"`
	Tag   string `json:"tag"`
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			fmt.Print(err)
			element.Field = err.Field()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}
