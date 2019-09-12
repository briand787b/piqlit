package obj

import (
	"context"
	"io"
)

// ObjectStore is anything capable of storing, retrieving, and deleting objects
type ObjectStore interface {
	ListNames(ctx context.Context) ([]string, error)
	Get(ctx context.Context, path string) (io.ReadCloser, error)
	Put(ctx context.Context, path string, content io.ReadCloser) error
	Delete(ctx context.Context, path string) error
}
