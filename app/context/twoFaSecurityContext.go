package context

type TwoFaSecurityContext struct {
	SecretKey string
	SecretUrl string
	CodeName  string
	QrBase64  string
}
