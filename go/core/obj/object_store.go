package obj

import (
	"context"
	"io"
)

// ObjectStore is anything capable of storing, retrieving, and deleting objects
type ObjectStore interface {
	Get(ctx context.Context, path string) (io.ReadCloser, error)
	Put(ctx context.Context, path string, content io.Reader) error
	Delete(ctx context.Context, path string) error
}
