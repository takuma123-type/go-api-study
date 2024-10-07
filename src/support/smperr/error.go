package smperr

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type appErr struct {
	code  int
	msg   string
	trace error
}

type AppError interface {
	Code() int
	Msg() string
	Trace() error
	Error() string
}

type BadRequestErr struct {
	*appErr
}
type InternalErr struct {
	*appErr
}
type NotFoundErr struct {
	*appErr
}

func (e *appErr) Code() int {
	return e.code
}
func (e *appErr) Msg() string {
	return e.msg
}
func (e *appErr) Error() string {
	return e.msg
}
func (e *appErr) Trace() error {
	return e.trace
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

var (
	ErrUserNotFound = &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   "user not found",
			trace: nil,
		},
	}
)

func BadRequest(msg string) *BadRequestErr {
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func BadRequestf(format string, args ...interface{}) *BadRequestErr {
	msg := fmt.Sprintf(format, args...)
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Internal(msg string) *InternalErr {
	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func NotFound(msg string) *NotFoundErr {
	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

type DatabaseError struct {
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database operation '%s' failed: %v", e.Operation, e.Err)
}

type JSONBindingError struct {
	Detail string
}

func (e *JSONBindingError) Error() string {
	return fmt.Sprintf("Failed to bind JSON: %s", e.Detail)
}

func HandleError(ctx *gin.Context, err error, statusCode int) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
