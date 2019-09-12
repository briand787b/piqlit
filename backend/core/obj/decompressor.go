package obj

import (
	"compress/gzip"
	"io"
)

type decompressor struct {
	rc io.ReadCloser
	gr *gzip.Reader
	pr *io.PipeReader
	pw *io.PipeWriter
}

func newDecompressor(rc io.ReadCloser) (*decompressor, error) {
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

	return &decompressor{
		gr: gr,
		pr: pr,
		pw: pw,
		rc: rc,
	}, nil
}

func (d *decompressor) Read(p []byte) (int, error) {
	return d.pr.Read(p)
}

func (d *decompressor) Close() error {
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
