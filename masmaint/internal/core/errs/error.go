package errs

import (
	"fmt"
)


/////////////////////////////////////////////////////////////////////////
type BadRequestError struct {
	Field string
	Type string
}

func NewBadRequestError(field, expectType string) error {
	return BadRequestError{Field: field, Type: expectType}
}

func (e BadRequestError) Error() string {
	if e.Field == "" {
		return "error: The content of the request is invalid."
	}
	return fmt.Sprintf("error: Field '%s' should be of type %s", e.Field, e.Type)
}

/////////////////////////////////////////////////////////////////////////
type UniqueConstraintError struct {
	Column string
}

func NewUniqueConstraintError(column string) error {
	return UniqueConstraintError{Column: column}
}

func (e UniqueConstraintError) Error() string {
	return fmt.Sprintf("error: Column '%s' should be unique.", e.Column)
}

/////////////////////////////////////////////////////////////////////////
type NotFoundError struct {}

func NewNotFoundError() error {
	return NotFoundError{}
}

func (e NotFoundError) Error() string {
	return "error: Not found"
}

/////////////////////////////////////////////////////////////////////////
type UnexpectedError struct {
	Message string
}

func NewUnexpectedError(message string) error {
	return UnexpectedError{Message: message}
}

func (e UnexpectedError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}