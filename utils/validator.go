package utils

import "github.com/go-playground/validator/v10"

func ValidateInput(input interface{}) error {
	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return err
	}

	return nil
}
