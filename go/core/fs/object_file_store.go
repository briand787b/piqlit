package fs

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/briand787b/piqlit/core/plog"
	"github.com/pkg/errors"
)

// ObjectFileStore is a ObjdctStore backed by a filesystem
type ObjectFileStore struct {
	l       plog.Logger
	baseDir string
}

// NewObjectFileStore returns a new, instantiated ObjectFileStore
func NewObjectFileStore(l plog.Logger, dir string) *ObjectFileStore {
	return &ObjectFileStore{
		l:       l,
		baseDir: dir,
	}
}

// Get x
func (ofs *ObjectFileStore) Get(ctx context.Context, path string) (io.ReadCloser, error) {
	fd, err := os.Open(filepath.Join(ofs.baseDir, path))
	if err != nil {
		return nil, errors.Wrap(err, "could not open file to GET")
	}

	return fd, nil
}

// Put x
func (ofs *ObjectFileStore) Put(ctx context.Context, path string, content io.Reader) error {
	fd, err := os.Create(filepath.Join(ofs.baseDir, path))
	if err != nil {
		return errors.Wrap(err, "could not open file to PUT")
	}

	defer ofs.l.Close(ctx, fd)

	if _, err := io.Copy(fd, content); err != nil {
		return errors.Wrap(err, "could not copy upload contents to file")
	}

	return nil
}

// Delete x
func (ofs *ObjectFileStore) Delete(ctx context.Context, path string) error {
	return nil
}
