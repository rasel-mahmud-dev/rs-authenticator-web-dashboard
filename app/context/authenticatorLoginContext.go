package context

import "rs/auth/app/dto"

type AuthenticatorLoginContext struct {
	RequestBody dto.AuthenticatorLoginRequestBody
}
