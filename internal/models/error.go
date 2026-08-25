package models

import "fmt"

type ErrorType string

const (
	RegistryError ErrorType = "registry"
	WorkerError   ErrorType = "worker"
	DBError       ErrorType = "db"
	DefaultError  ErrorType = "default"
)

type InternalError struct {
	errorType ErrorType
	err       error
	msg       string
}

func NewInternalError(t ErrorType, msg string) *InternalError {
	return &InternalError{errorType: t, err: fmt.Errorf("[%s]: %s", t, msg), msg: msg}
}

func (err *InternalError) ErrorType() ErrorType {
	return err.errorType
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("[%s] %s", e.errorType, e.msg)
}

func (e *InternalError) Unwrap() error {
	return e.err
}
