package errors

import (
	"fmt"

	"github.com/Aboagye-Dacosta/shopBackend/internal/messages"
)

type AppError struct {
	Code        int
	Entity      string
	UserMessage string
	DevMessage  string
	Err         error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Entity, e.DevMessage)
}

func New(entity string, code int, err error) *AppError {
	msg := messages.Error(entity, code)
	return &AppError{
		Code:        code,
		Entity:      entity,
		UserMessage: msg,
		DevMessage:  messages.ErrorDev(entity, code),
		Err:         err,
	}
}
