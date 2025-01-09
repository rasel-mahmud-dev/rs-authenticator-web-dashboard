package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"rs/auth/app/dto"
)

var validate = validator.New()

func ValidateLoginRequest(loginRequest *dto.LoginRequest) error {
	err := validate.Struct(loginRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation failed for field '%s': %s", e.Field(), e.Tag())
		}
	}
	return nil
}
