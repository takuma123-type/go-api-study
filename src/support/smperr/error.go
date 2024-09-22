package smperr

import (
	"fmt"
	"net/http"
	"unicode/utf8"

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
type UnauthorizedErr struct {
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

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

const lastNameLength = 50

func validateLastName(last string) (interface{}, error) {
	if l := utf8.RuneCountInString(last); l > lastNameLength {
		return nil, fmt.Errorf("last name must be less than %d characters", lastNameLength)
	}
	return last, nil
}

const firstNameLength = 50

func validateFirstName(first string) (interface{}, error) {
	if l := utf8.RuneCountInString(first); l > firstNameLength {
		return nil, fmt.Errorf("first name must be less than %d characters", firstNameLength)
	}
	return first, nil
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

func BadRequestf(format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func BadRequestWrapf(err2 error, format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
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

func Internalf(format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func InternalWrapf(err2 error, format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
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

func NotFoundf(format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func NotFoundWrapf(err2 error, format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

func Unauthorized(msg string) *UnauthorizedErr {
	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Unauthorizedf(format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}
func UnauthorizedWrapf(err2 error, format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	// 1.20からWrapがJoinに変わる
	//err = errors.Join(err, err2)
	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

type UserNotFoundError struct {
	ID string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with ID %s not found", e.ID)
}

// UpdateFailedError is returned when updating the user fails.
type UpdateFailedError struct {
	Reason string
}

func (e *UpdateFailedError) Error() string {
	return fmt.Sprintf("Failed to update user: %s", e.Reason)
}

// DatabaseError represents a generic database operation failure.
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

type UpdateUserError struct {
	Reason string
}

func (e *UpdateUserError) Error() string {
	return fmt.Sprintf("Failed to update user: %s", e.Reason)
}

type InvalidURIError struct {
	Detail string
}

func (e *InvalidURIError) Error() string {
	return fmt.Sprintf("Invalid URI: %s", e.Detail)
}

type InvalidJSONError struct {
	Detail string
}

func (e *InvalidJSONError) Error() string {
	return fmt.Sprintf("Invalid JSON: %s", e.Detail)
}

type DatabaseConnectionError struct {
	Reason string
}

func (e *DatabaseConnectionError) Error() string {
	return fmt.Sprintf("Failed to connect to the database: %s", e.Reason)
}
