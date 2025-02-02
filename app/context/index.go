package LoginContext

import "context"

type LoginContext struct {
	context.Context
}
type AccessToken string

type CustomKey string

const TokenKey CustomKey = "token"

func (c *LoginContext) SetAccessToken(token AccessToken) context.Context {
	return context.WithValue(c, TokenKey, token)
}

func (c *LoginContext) GetAccessToken() (AccessToken, bool) {
	token, ok := c.Value(TokenKey).(AccessToken)
	return token, ok
}
