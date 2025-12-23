package error

import "net/http"

type HttpError struct {
	Status  int
	Code    string
	Message string
}

func Map(err error) HttpError {
	if appErr, ok := err.(*AppError); ok {
		return HttpError{
			Status:  appErr.Status,
			Code:    appErr.Code,
			Message: appErr.Message,
		}
	}

	return HttpError{
		Status:  http.StatusInternalServerError,
		Code:    ErrInternal,
		Message: "internal server error",
	}
}
