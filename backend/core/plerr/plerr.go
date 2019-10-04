package plerr

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// // ValidationErr is when the attempted operation fails validation
// // due to invalid user input
// type ValidationErr error

var (
	// ErrValidation is when validation of the resource failed
	ErrValidation = errors.New("resource is invalid")

	// ErrNotFound x
	ErrNotFound = errors.New("resource could not be found")

	// ErrUnauthorized x
	ErrUnauthorized = errors.New("authorization failed")

	ErrServerErr = errors.New("internal server error")
)

func GetExternalMsg(e error) string {
	if e == nil {
		return ""
	}

	switch c := errors.Cause(e); {
	case c == ErrValidation:
		if es := strings.Split(e.Error(), ":"); len(es) > 1 {
			return fmt.Sprintf("%s: %s", strings.TrimSpace(es[len(es)-1]), es[len(es)-2])
		}

		fallthrough
	default:
		return c.Error()
	}
}

func GetInternalMsg(e error) string {
	if e == nil {
		return ""
	}

	return e.Error()
}

func NewErrValidation(reasonMsg string) error {
	err := errors.Wrap(ErrValidation, reasonMsg)
	return err
}

func NewErrNotFound(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrNotFound()`")
	}

	return errors.Wrap(ErrNotFound, e.Error())
}
