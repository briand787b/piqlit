package obj

import (
	"compress/gzip"
	"context"
	"io"

	"github.com/briand787b/piqlit/core/plog"

	"github.com/pkg/errors"
)

// GZCompressor wraps uncompressed data and provides it in a compressed format
// through its Read method
type GZCompressor struct {
	ctx context.Context // yes this is bad, but i cant pass a ctx to a func that conforms to a stdlib interface
	l   plog.Logger
	r   io.Reader
	gw  *gzip.Writer
	pw  *io.PipeWriter
	pr  *io.PipeReader
}

// NewGZCompressor instantiates a GZCompressor
func NewGZCompressor(ctx context.Context, l plog.Logger, r io.Reader) *GZCompressor {
	pr, pw := io.Pipe()
	gw := gzip.NewWriter(pw)

	go func() {
		if _, err := io.Copy(gw, r); err != nil {
			l.Error(ctx, "NewGZCompressor could not copy from ReadCloser", "error", err)
			l.Close(ctx, gw)
			pw.CloseWithError(err)
		} else {
			l.Close(ctx, gw)
			l.Close(ctx, pw)
		}
	}()

	return &GZCompressor{
		ctx: ctx,
		l:   l,
		r:   r,
		gw:  gw,
		pw:  pw,
		pr:  pr,
	}
}

// Read reads compressed data from GZCompressor
func (c *GZCompressor) Read(p []byte) (int, error) {
	return c.pr.Read(p)
}

// Close closes all the resoprces used by GZCompressor
func (c *GZCompressor) Close() error {
	var returnErr error

	if err := c.pr.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		c.l.Error(c.ctx, "could not close pipe reader", "error", err)
	}

	return returnErr
}
