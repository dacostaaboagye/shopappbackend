package errors

import (
	"database/sql"
	stdErrors "errors"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type DbErrorCode int

const (
	DbErrNotFound DbErrorCode = iota
	DbErrConflict
	DbErrBadRequest   // validation errors, constraint violations, null violations
	DbErrUnauthorized // permission denied
	DbErrForbidden    // insufficient privileges
	DbErrTimeout      // query timeout, statement timeout
	DbErrConnection   // connection pool exhausted, DB unavailable
)

// DbError represents a database-layer error with a stable code and optional context.
type DbError struct {
	Code DbErrorCode
	Op   string // operation, e.g., "userRepo.FindByEmail"
	Err  error  // wrapped driver error (optional)
}

func (e *DbError) Error() string {
	if e.Op != "" {
		return e.Op
	}
	return "db error"
}

func (e *DbError) Unwrap() error { return e.Err }

// Is allows errors.Is(err, ErrDbNotFound) style checks by matching the Code.
func (e *DbError) Is(target error) bool {
	t, ok := target.(*DbError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// Sentinels for code-based comparisons.
var (
	ErrDbNotFound     = &DbError{Code: DbErrNotFound}
	ErrDbConflict     = &DbError{Code: DbErrConflict}
	ErrDbBadRequest   = &DbError{Code: DbErrBadRequest}
	ErrDbUnauthorized = &DbError{Code: DbErrUnauthorized}
	ErrDbForbidden    = &DbError{Code: DbErrForbidden}
	ErrDbTimeout      = &DbError{Code: DbErrTimeout}
	ErrDbConnection   = &DbError{Code: DbErrConnection}
)

// Constructors to create typed DB errors with context.
func NewDbNotFound(op string, cause error) *DbError {
	return &DbError{Code: DbErrNotFound, Op: op, Err: cause}
}

func NewDbConflict(op string, cause error) *DbError {
	return &DbError{Code: DbErrConflict, Op: op, Err: cause}
}

func NewDbBadRequest(op string, cause error) *DbError {
	return &DbError{Code: DbErrBadRequest, Op: op, Err: cause}
}

func NewDbUnauthorized(op string, cause error) *DbError {
	return &DbError{Code: DbErrUnauthorized, Op: op, Err: cause}
}

func NewDbForbidden(op string, cause error) *DbError {
	return &DbError{Code: DbErrForbidden, Op: op, Err: cause}
}

func NewDbTimeout(op string, cause error) *DbError {
	return &DbError{Code: DbErrTimeout, Op: op, Err: cause}
}

func NewDbConnection(op string, cause error) *DbError {
	return &DbError{Code: DbErrConnection, Op: op, Err: cause}
}

// MapDbErrToHTTP maps a DbError to an HTTP status code.
func MapDbErrToHTTP(err error) int {
	var dbe *DbError
	if stdErrors.As(err, &dbe) {
		switch dbe.Code {
		case DbErrNotFound:
			return http.StatusNotFound
		case DbErrConflict:
			return http.StatusConflict
		case DbErrBadRequest:
			return http.StatusBadRequest
		case DbErrUnauthorized:
			return http.StatusUnauthorized
		case DbErrForbidden:
			return http.StatusForbidden
		case DbErrTimeout:
			return http.StatusRequestTimeout
		case DbErrConnection:
			return http.StatusServiceUnavailable
		}
	}
	return http.StatusInternalServerError
}

// ClassifyDbError analyzes a database error (from GORM, sql, or PostgreSQL) and wraps it as a DbError.
// This handles common PostgreSQL errors automatically.
func ClassifyDbError(op string, err error) error {
	if err == nil {
		return nil
	}

	// Check if already a DbError
	if dbe, ok := err.(*DbError); ok {
		return dbe
	}

	errStr := strings.ToLower(err.Error())

	// Handle GORM errors
	if stdErrors.Is(err, gorm.ErrRecordNotFound) || stdErrors.Is(err, sql.ErrNoRows) {
		return NewDbNotFound(op, err)
	}

	// PostgreSQL/GORM constraint violations and common errors
	// Unique constraint violations (23505)
	if strings.Contains(errStr, "duplicate key") ||
		strings.Contains(errStr, "unique constraint") ||
		strings.Contains(errStr, "violates unique constraint") ||
		strings.Contains(errStr, "23505") {
		return NewDbConflict(op, err)
	}

	// Foreign key violations (23503)
	if strings.Contains(errStr, "foreign key constraint") ||
		strings.Contains(errStr, "violates foreign key constraint") ||
		strings.Contains(errStr, "23503") {
		return NewDbBadRequest(op, err)
	}

	// Not null violations (23502)
	if strings.Contains(errStr, "null value") ||
		strings.Contains(errStr, "violates not-null constraint") ||
		strings.Contains(errStr, "23502") {
		return NewDbBadRequest(op, err)
	}

	// Check constraint violations (23514)
	if strings.Contains(errStr, "check constraint") ||
		strings.Contains(errStr, "23514") {
		return NewDbBadRequest(op, err)
	}

	// Permission/authorization errors
	if strings.Contains(errStr, "permission denied") ||
		strings.Contains(errStr, "insufficient privilege") ||
		strings.Contains(errStr, "42501") {
		return NewDbForbidden(op, err)
	}

	// Connection errors
	if strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "no connection available") ||
		strings.Contains(errStr, "connection pool") ||
		strings.Contains(errStr, "too many connections") ||
		strings.Contains(errStr, "08000") || // Connection exception
		strings.Contains(errStr, "08003") || // Connection does not exist
		strings.Contains(errStr, "08006") { // Connection failure
		return NewDbConnection(op, err)
	}

	// Timeout errors
	if strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded") ||
		strings.Contains(errStr, "57014") { // Query canceled
		return NewDbTimeout(op, err)
	}

	// Default: treat as internal server error (don't wrap, let it bubble up)
	return err
}

// FromDb converts a DB-layer error into an AppError using the entity's message registry.
// It automatically classifies the error if it's not already a DbError.
func FromDb(entity string, err error) *AppError {
	if err == nil {
		return nil
	}
	classified := ClassifyDbError("", err)
	return New(entity, MapDbErrToHTTP(classified), classified)
}
