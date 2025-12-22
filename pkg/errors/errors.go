package errors

import "fmt"

type ErrorType string

const (
	ValidationError ErrorType = "validation"
	ExecutionError  ErrorType = "execution"
	ConfigError     ErrorType = "config"
	UnknownError    ErrorType = "unknown"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

func NewValidationError(msg string, err error) *AppError {
	return &AppError{Type: ValidationError, Message: msg, Err: err}
}

func NewExecutionError(msg string, err error) *AppError {
	return &AppError{Type: ExecutionError, Message: msg, Err: err}
}

func NewConfigError(msg string, err error) *AppError {
	return &AppError{Type: ConfigError, Message: msg, Err: err}
}

func NewUnknownError(msg string, err error) *AppError {
	return &AppError{Type: UnknownError, Message: msg, Err: err}
}
