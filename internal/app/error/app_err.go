package error

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

func New(
	code string,
	message string,
	status int,
	err error,
) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
		Err:     err,
	}
}
