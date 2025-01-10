package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct[T any](payload *T) error {
	err := validate.Struct(payload)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation failed for field '%s': %s", e.Field(), e.Tag())
		}
	}
	return nil
}
