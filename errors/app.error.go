package errors

import (
	"net/http"
)

type AppError struct {
	Code    int
	Type    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Type:    "Not Found",
		Message: message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Type:    "Bad Request",
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Type:    "Validation Error",
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Type:    "Internal Server Error",
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Type:    "Unauthorized",
		Message: message,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Type:    "Forbidden",
		Message: message,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    http.StatusConflict,
		Type:    "Conflict",
		Message: message,
	}
}

func NewTooManyRequestsError(message string) *AppError {
	return &AppError{
		Code:    http.StatusTooManyRequests,
		Type:    "Too Many Requests",
		Message: message,
	}
}
