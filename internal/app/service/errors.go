package service

import "net/http"

type AppError struct {
	Code    string
	Message string
	Status  int
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(code, message string, status int, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
		Err:     err,
	}
}

var (
	ErrUnauthorized = NewAppError(
		"UNAUTHORIZED",
		"unauthorized",
		http.StatusUnauthorized,
		nil,
	)

	ErrInvalidCredential = NewAppError(
		"INVALID_CREDENTIAL",
		"email or password is incorrect",
		http.StatusUnauthorized,
		nil,
	)

	ErrEmailAlreadyUsed = NewAppError(
		"EMAIL_ALREADY_USED",
		"email already registered",
		http.StatusBadRequest,
		nil,
	)
)
