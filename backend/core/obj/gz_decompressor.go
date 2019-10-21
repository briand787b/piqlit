package obj

import (
	"compress/gzip"
	"io"
)

// GZDecompressor is capable of reading uncompressed data
// and having compressed data read out of it
type GZDecompressor struct {
	rc io.ReadCloser
	gr *gzip.Reader
	pr *io.PipeReader
	pw *io.PipeWriter
}

// NewGZDecompressor returns a new GZDecompressor
func NewGZDecompressor(rc io.ReadCloser) (*GZDecompressor, error) {
	gr, err := gzip.NewReader(rc)
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	go func() {
		if _, err := io.Copy(pw, gr); err != nil {
			pw.CloseWithError(err)
		} else {
			pw.Close()
		}
	}()

	return &GZDecompressor{
		gr: gr,
		pr: pr,
		pw: pw,
		rc: rc,
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
	if err := d.rc.Close(); err != nil {
		return err
	}

	if err := d.gr.Close(); err != nil {
		return err
	}

	if err := d.pw.Close(); err != nil {
		return err
	}

	if err := d.pr.Close(); err != nil {
		return err
	}

	return nil
}
