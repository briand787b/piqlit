package obj

import (
	"compress/gzip"
	"context"
	"io"

	"github.com/briand787b/piqlit/core/plog"

	"github.com/pkg/errors"
)

// GZDecompressor is capable of reading uncompressed data
// and having compressed data read out of it
type GZDecompressor struct {
	ctx context.Context // yes this is bad, but i cant pass a ctx to a func that conforms to a stdlib interface
	l   plog.Logger
	rc  io.ReadCloser
	gr  *gzip.Reader
	pr  *io.PipeReader
	pw  *io.PipeWriter
}

// NewGZDecompressor returns a new GZDecompressor
func NewGZDecompressor(ctx context.Context, l plog.Logger, rc io.ReadCloser) (*GZDecompressor, error) {
	gr, err := gzip.NewReader(rc)
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	go func() {
		if _, err := io.Copy(pw, gr); err != nil {
			l.Error(ctx, "could not copy gzip reader to pipe writer", "error", err.Error())
			l.Close(ctx, rc)
			pw.CloseWithError(err)
		} else {
			l.Close(ctx, rc)
			l.Close(ctx, pw)
		}
	}()

	return &GZDecompressor{
		ctx: ctx,
		l:   l,
		gr:  gr,
		pr:  pr,
		pw:  pw,
		rc:  rc,
	}, nil
}

// Read reads from the uncompressed stream and fills in the
// supplied byte slice with compressed data
func (d *GZDecompressor) Read(p []byte) (int, error) {
	return d.pr.Read(p)
}

// Close closes all the associated resources of the
// GZDecompressor
func (d *GZDecompressor) Close() error {
	var returnErr error

	if err := d.gr.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		d.l.Error(d.ctx, "could not close gzip reader", "error", err)
	}

	if err := d.pr.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		d.l.Error(d.ctx, "could not close pipe reader", "error", err)
	}

	return returnErr
}
