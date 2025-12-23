package error

const (
	// auth
	ErrUnauthorized      = "AUTH_UNAUTHORIZED"
	ErrInvalidCredential = "AUTH_INVALID_CREDENTIAL"
	ErrTokenExpired      = "AUTH_TOKEN_EXPIRED"

	// validation
	ErrValidation = "VALIDATION_ERROR"

	// system
	ErrInternal = "INTERNAL_ERROR"
	ErrDatabase = "DATABASE_ERROR"
)
