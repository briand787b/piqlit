package obj

import (
	"compress/gzip"
	"io"
	"log"

	"github.com/pkg/errors"
)

// GZCompressor wraps uncompressed data and provides it in a compressed format
// through its Read method
type GZCompressor struct {
	rc io.ReadCloser
	gw *gzip.Writer
	pw *io.PipeWriter
	pr *io.PipeReader
}

// NewGZCompressor instantiates a GZCompressor
func NewGZCompressor(rc io.ReadCloser) *GZCompressor {
	pr, pw := io.Pipe()
	gw := gzip.NewWriter(pw)

	go func() {
		if _, err := io.Copy(gw, rc); err != nil {
			log.Println("NewGZCompressor: could not copy from ReadCloser: ", err)
			log.Println(gw.Close())
			pw.CloseWithError(err)
		} else {
			log.Println(gw.Close())
			log.Println(pw.Close())
		}
	}()

	return &GZCompressor{
		rc: rc,
		gw: gw,
		pw: pw,
		pr: pr,
	}
}

// Read reads compressed data from GZCompressor
func (c *GZCompressor) Read(p []byte) (int, error) {
	return c.pr.Read(p)
}

// Close closes all the resoprces used by GZCompressor
func (c *GZCompressor) Close() error {
	var returnErr error

	if err := c.rc.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		log.Println("could not close rc: ", err)
	}
	if err := c.gw.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		log.Println("could not close gw: ", err)
	}
	if err := c.pw.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		log.Println("Closecould not close pw: ", err)
	}
	if err := c.pr.Close(); err != nil {
		returnErr = errors.Wrap(returnErr, err.Error())
		log.Println("could not close pr: ", err)
	}

	return returnErr
}
