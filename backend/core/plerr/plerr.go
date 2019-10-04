package plerr

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrValidation is when validation of the resource failed
	ErrValidation = errors.New("resource is invalid")

	// ErrNotFound is when the request resource does not exist
	ErrNotFound = errors.New("resource could not be found")

	// ErrUnauthorized is when the requestor is unauthorized to perform
	// the requested action
	ErrUnauthorized = errors.New("authorization failed")

	// ErrInternal is when an error results from internal software failures
	ErrInternal = errors.New("internal server error")
)

// GetExternalMsg extracts the message for the error that is suitable
// for displaying externally
func GetExternalMsg(e error) string {
	if e == nil {
		return ""
	}

	switch c := errors.Cause(e); {
	case c == ErrValidation:
		if es := strings.Split(e.Error(), ":"); len(es) > 1 {
			return strings.TrimSpace(fmt.Sprintf("%s: %s", es[len(es)-1], es[len(es)-2]))
		}

		fallthrough
	default:
		return c.Error()
	}
}

// GetInternalMsg extracts the message for the error that is suitable
// for internal logging
func GetInternalMsg(e error) string {
	if e == nil {
		return ""
	}

	return e.Error()
}

// NewErrValidation returns a wrapped ErrValidation
func NewErrValidation(reasonMsg string) error {
	err := errors.Wrap(ErrValidation, reasonMsg)
	return err
}

// NewErrNotFound returns a wrapped ErrNotFound
func NewErrNotFound(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrNotFound()`")
	}

	return errors.Wrap(ErrNotFound, e.Error())
}
