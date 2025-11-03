package errors

import (
	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
)

// Authorization error codes (re-exported for application use)
const (
	INVALID_TOKEN     = codes.INVALID_TOKEN
	EXPIRED_TOKEN     = codes.EXPIRED_TOKEN
	NO_TOKEN_PROVIDED = codes.NO_TOKEN_PROVIDED
)

// NewAuth creates a new AppError for authorization using errors.New under the hood.
func NewAuth(code int, err error) *AppError {
	return New(entities.AUTHORIZATION, code, err)
}

// InvalidToken returns an AppError for invalid token scenarios.
func InvalidToken(err error) *AppError {
	return NewAuth(INVALID_TOKEN, err)
}

// ExpiredToken returns an AppError for expired token scenarios.
func ExpiredToken(err error) *AppError {
	return NewAuth(EXPIRED_TOKEN, err)
}

// NoTokenProvided returns an AppError when the Authorization token is missing.
func NoTokenProvided(err error) *AppError {
	return NewAuth(NO_TOKEN_PROVIDED, err)
}
