package controller

import (
	"context"
	"net/http"

	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/pkg/errors"

	"github.com/go-chi/render"
)

// ErrResponse is the error response that gets sent back in the response
type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

func newErrResponse(ctx context.Context, l plog.Logger, err error) *ErrResponse {
	er := ErrResponse{
		StatusText: perr.GetExternalMsg(ctx, l, err),
	}

	switch errors.Cause(err) {
	case perr.ErrInvalid:
		er.HTTPStatusCode = http.StatusBadRequest
	case perr.ErrNotFound:
		er.HTTPStatusCode = http.StatusNotFound
	case perr.ErrUnauthorized:
		er.HTTPStatusCode = http.StatusUnauthorized
	default:
		er.HTTPStatusCode = http.StatusInternalServerError
	}

	l.Error(ctx, "error handling request",
		"status_code", er.HTTPStatusCode,
		"error", err.Error(),
	)
	return &er
}

// Render allows ErrResponse to satisfy render.Renderer interface
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
