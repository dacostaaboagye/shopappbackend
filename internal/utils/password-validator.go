package utils

import (
    "errors"
    "regexp"
)

var (
    ErrTooShort       = errors.New("password must be at least 8 characters long")
    ErrMissingUpper   = errors.New("password must contain at least one uppercase letter")
    ErrMissingLower   = errors.New("password must contain at least one lowercase letter")
    ErrMissingNumber  = errors.New("password must contain at least one number")
    ErrMissingSpecial = errors.New("password must contain at least one special character")
)

// ValidatePassword checks a password against common policy rules
func ValidatePassword(password string) error {
    if len(password) < 8 {
        return ErrTooShort
    }

    upper := regexp.MustCompile(`[A-Z]`)
    lower := regexp.MustCompile(`[a-z]`)
    number := regexp.MustCompile(`[0-9]`)
    special := regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?/{}\-]`)

    switch {
    case !upper.MatchString(password):
        return ErrMissingUpper
    case !lower.MatchString(password):
        return ErrMissingLower
    case !number.MatchString(password):
        return ErrMissingNumber
    case !special.MatchString(password):
        return ErrMissingSpecial
    }

    return nil
}
