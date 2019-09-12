package obj

import (
	"context"
	"io"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// ObjectFileStore is a ObjdctStore backed by a filesystem
type ObjectFileStore struct {
	l       *logrus.Logger
	pattern string
}

// NewObjectFileStore returns a new, instantiated ObjectFileStore
func NewObjectFileStore(l *logrus.Logger, dir string) *ObjectFileStore {
	return &ObjectFileStore{
		l:       l,
		pattern: filepath.Join(dir, "*"),
	}
}

// ListNames x
func (ofs *ObjectFileStore) ListNames(ctx context.Context) ([]string, error) {
	ofs.l.WithFields(logrus.Fields{
		"pattern": ofs.pattern,
	}).Info("ObjectFileStore.ListNames()")
	fs, err := filepath.Glob(ofs.pattern)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(fs); i++ {
		fs[i] = filepath.Base(fs[i])
	}

	return fs, nil
}

// Get x
func (ofs *ObjectFileStore) Get(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, nil
}

// Put x
func (ofs *ObjectFileStore) Put(ctx context.Context, path string, content io.ReadCloser) error {
	return nil
}

// Delete x
func (ofs *ObjectFileStore) Delete(ctx context.Context, path string) error {
	return nil
}
