package statusCode

var HttpStatus map[StatusCodeType]int

type StatusCodeType string

var (
	OK                             StatusCodeType = "SUCCESS-000"
	LOGIN_SUCCESS                  StatusCodeType = "SUCCESS-001"
	LOGOUT_SUCCESS                 StatusCodeType = "SUCCESS-002"
	ACCOUNT_CREATED                StatusCodeType = "SUCCESS-003"
	PASSWORD_CHANGED               StatusCodeType = "SUCCESS-004"
	TWO_FACTOR_ENABLED             StatusCodeType = "SUCCESS-005"
	TWO_FACTOR_VERIFIED            StatusCodeType = "SUCCESS-006"
	EMAIL_VERIFIED                 StatusCodeType = "SUCCESS-007"
	PHONE_VERIFIED                 StatusCodeType = "SUCCESS-008"
	PROFILE_UPDATED                StatusCodeType = "SUCCESS-009"
	SESSION_EXTENDED               StatusCodeType = "SUCCESS-010"
	DEVICE_ADDED                   StatusCodeType = "SUCCESS-011"
	DEVICE_VERIFIED                StatusCodeType = "SUCCESS-012"
	TOKEN_REFRESHED                StatusCodeType = "SUCCESS-013"
	PERMISSIONS_UPDATED            StatusCodeType = "SUCCESS-014"
	CAPTCHA_VERIFIED               StatusCodeType = "SUCCESS-015"
	ACCOUNT_REACTIVATED            StatusCodeType = "SUCCESS-016"
	PASSWORD_RESET_INITIATED       StatusCodeType = "SUCCESS-017"
	PASSWORD_RESET_COMPLETED       StatusCodeType = "SUCCESS-018"
	ADMIN_LOGIN_SUCCESS            StatusCodeType = "SUCCESS-019"
	AUTH_SERVICE_READY             StatusCodeType = "SUCCESS-020"
	PASSWORD_EXPIRING_SOON         StatusCodeType = "WARN-001"
	SESSION_EXPIRING_SOON          StatusCodeType = "WARN-002"
	UNUSUAL_LOGIN_LOCATION         StatusCodeType = "WARN-003"
	UNUSUAL_LOGIN_DEVICE           StatusCodeType = "WARN-004"
	MULTIPLE_FAILED_ATTEMPTS       StatusCodeType = "WARN-005"
	TWO_FACTOR_NOT_ENABLED         StatusCodeType = "WARN-006"
	OLD_PASSWORD_REUSED            StatusCodeType = "WARN-007"
	ACCOUNT_NOT_VERIFIED           StatusCodeType = "WARN-008"
	INSECURE_CONNECTION            StatusCodeType = "WARN-009"
	SESSION_ON_MULTIPLE_DEVICES    StatusCodeType = "WARN-014"
	UNVERIFIED_DEVICE              StatusCodeType = "WARN-015"
	UNVERIFIED_IP                  StatusCodeType = "WARN-016"
	TOKEN_ALMOST_EXPIRED           StatusCodeType = "WARN-018"
	RECENT_PASSWORD_CHANGE         StatusCodeType = "WARN-019"
	ADMIN_OVERRIDE_LOGIN           StatusCodeType = "WARN-020"
	INVALID_CREDENTIALS            StatusCodeType = "ERR-001"
	USER_NOT_FOUND                 StatusCodeType = "ERR-002"
	ACCOUNT_LOCKED                 StatusCodeType = "ERR-003"
	TOKEN_EXPIRED                  StatusCodeType = "ERR-004"
	TOKEN_INVALID                  StatusCodeType = "ERR-005"
	PERMISSION_DENIED              StatusCodeType = "ERR-006"
	SESSION_EXPIRED                StatusCodeType = "ERR-007"
	ACCOUNT_DISABLED               StatusCodeType = "ERR-008"
	PASSWORD_EXPIRED               StatusCodeType = "ERR-009"
	PASSWORD_POLICY_VIOLATION      StatusCodeType = "ERR-010"
	AUTHENTICATION_TIMEOUT         StatusCodeType = "ERR-011"
	UNAUTHORIZED_ACCESS            StatusCodeType = "ERR-012"
	TWO_FACTOR_REQUIRED            StatusCodeType = "ERR-013"
	TWO_FACTOR_FAILED              StatusCodeType = "ERR-014"
	INVALID_OTP                    StatusCodeType = "ERR-015"
	OTP_EXPIRED                    StatusCodeType = "ERR-016"
	IP_BLOCKED                     StatusCodeType = "ERR-017"
	DEVICE_NOT_RECOGNIZED          StatusCodeType = "ERR-018"
	CAPTCHA_FAILED                 StatusCodeType = "ERR-019"
	AUTH_SERVICE_UNAVAILABLE       StatusCodeType = "ERR-020"
	ACCOUNT_CREATION_FAILED        StatusCodeType = "ERR-021"
	INVALID_JSON_FORMAT            StatusCodeType = "ERR-022"
	REQUEST_VALIDATION_FAILED      StatusCodeType = "ERR-023"
	INTERNAL_SERVER_ERROR          StatusCodeType = "ERR-024"
	DUPLICATE_ENTITY               StatusCodeType = "ERR-025"
	PASSWORD_NOT_CONFIGURED        StatusCodeType = "ERR-026"
	UNAUTHORIZED                   StatusCodeType = "ERR-027"
	ACCESS_TOKEN_MISSED            StatusCodeType = "ERR-028"
	UNAUTHORIZED_SESSION_MISMATCH  StatusCodeType = "ERR-028"
	UNAUTHORIZED_SESSION_REVOKED   StatusCodeType = "ERR-028"
	UNAUTHORIZED_SESSION_NOT_FOUND StatusCodeType = "ERR-028"
	UNAUTHORIZED_TOKEN_INVALID     StatusCodeType = "ERR-028"
	UNAUTHORIZED_TOKEN_MISSING     StatusCodeType = "ERR-028"
)

func init() {
	HttpStatus = map[StatusCodeType]int{
		INVALID_CREDENTIALS:       400,
		USER_NOT_FOUND:            404,
		ACCOUNT_LOCKED:            423,
		TOKEN_EXPIRED:             401,
		TOKEN_INVALID:             400,
		PERMISSION_DENIED:         403,
		SESSION_EXPIRED:           401,
		ACCOUNT_DISABLED:          403,
		PASSWORD_EXPIRED:          401,
		PASSWORD_POLICY_VIOLATION: 400,
		AUTHENTICATION_TIMEOUT:    408,
		UNAUTHORIZED_ACCESS:       401,
		TWO_FACTOR_REQUIRED:       401,
		TWO_FACTOR_FAILED:         401,
		INVALID_OTP:               400,
		OTP_EXPIRED:               400,
		IP_BLOCKED:                403,
		DEVICE_NOT_RECOGNIZED:     401,
		CAPTCHA_FAILED:            400,
		AUTH_SERVICE_UNAVAILABLE:  503,
	}
}
