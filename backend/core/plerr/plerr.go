package plerr

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrInvalid is when validation of the resource failed
	ErrInvalid = Error("request invalid")

	// ErrNotFound is when the request resource does not exist
	ErrNotFound = Error("resource could not be found")

	// ErrUnauthorized is when the requestor is unauthorized to perform
	// the requested action
	ErrUnauthorized = Error("authorization failed")

	// ErrInternal is when an error results from internal software failures
	ErrInternal = Error("internal server error")
)

// GetExternalMsg extracts the message for the error that is suitable
// for displaying externally
func GetExternalMsg(e error) string {
	if e == nil {
		return ""
	}

	switch c := errors.Cause(e); {
	case c == ErrInvalid:
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

// NewErrInvalid returns a wrapped ErrInvalid
func NewErrInvalid(reasonMsg string) error {
	err := errors.Wrap(ErrInvalid, reasonMsg)
	return err
}

// NewErrNotFound returns a wrapped ErrNotFound
func NewErrNotFound(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrNotFound()`")
	}

	return errors.Wrap(ErrNotFound, e.Error())
}

// NewErrInternal returns a wrapped ErrNewInternal
func NewErrInternal(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrInternal()`")
	}

	return errors.Wrap(ErrInternal, e.Error())
}
