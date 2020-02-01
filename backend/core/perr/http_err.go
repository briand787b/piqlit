package perr

import (
	"context"
	"net/http"

	"github.com/briand787b/piqlit/core/plog"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// HTTPError is an HTTPError
type HTTPError struct {
	Message string

	err error
	l   plog.Logger
	ctx context.Context
}

// NewHTTPErrorFromError returns a new HTTPError from an existing error
func NewHTTPErrorFromError(ctx context.Context, e error, msg string, l plog.Logger) *HTTPError {
	return &HTTPError{
		err: errors.Wrap(e, msg),
		l:   l,
		ctx: ctx,
	}

}

// NewValidationHTTPErrorFromError returns a new HTTPError of type Validation from an already existing error
func NewValidationHTTPErrorFromError(ctx context.Context, e error, msg string, l plog.Logger) *HTTPError {
	return &HTTPError{
		err: NewErrInvalid(msg),
		l:   l,
		ctx: ctx,
	}
}

// NewNotFoundHTTPError instantiates a new HTTPError of type NotFound
func NewNotFoundHTTPError(ctx context.Context, msg string, l plog.Logger) *HTTPError {
	return &HTTPError{
		err: NewErrNotFound(errors.New(msg)),
		l:   l,
		ctx: ctx,
	}
}

// NewInternalServerHTTPError returns a new HTTPError of type InternalServer
func NewInternalServerHTTPError(ctx context.Context, msg string, l plog.Logger) *HTTPError {
	return &HTTPError{
		err: NewErrInternal(errors.New(msg)),
		l:   l,
		ctx: ctx,
	}
}

// Render allows ErrResponse to satisfy render.Renderer interface
func (e *HTTPError) Render(w http.ResponseWriter, r *http.Request) error {
	var httpStatusCode int
	switch errors.Cause(e.err) {
	case ErrInvalid:
		httpStatusCode = http.StatusBadRequest
	case ErrNotFound:
		httpStatusCode = http.StatusNotFound
	case ErrUnauthorized:
		httpStatusCode = http.StatusUnauthorized
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	e.l.Error(e.ctx, "error handling request",
		"status_code", httpStatusCode,
		"error", e.err.Error(),
	)

	e.Message = GetExternalMsg(e.ctx, e.l, e.err)

	render.Status(r, httpStatusCode)
	return nil
}
