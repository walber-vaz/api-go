package logger_errors

import "net/http"

type LoggerErrors struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *LoggerErrors) Error() string {
	return e.Message
}

func NewLoggerErrors(message string, err string, code int, causes []Causes) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, err string) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *LoggerErrors {
	return &LoggerErrors{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
