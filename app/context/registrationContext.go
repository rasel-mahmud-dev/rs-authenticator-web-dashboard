package context

import "rs/auth/app/dto"

type RegistrationContext struct {
	Payload dto.RegisterRequestBody
}
