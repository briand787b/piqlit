package model

import (
	"time"

	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"
)

// Media is a container for viewable works of art, or for child media
type Media struct {
	ID           int              `db:"id"`
	Name         string           `db:"name"`
	Encoding     obj.Encoding     `db:"encoding"`
	UploadStatus obj.UploadStatus `db:"upload_status"`
	Children     []Media          `db:"children"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
}

// // Persist saves a Media to persistent storage
// func (m *Media) Persist(ctx context.Context, l plog.Logger) error {

// }

// Validate returns an error if the Media is not properly
// configured for persistent storage
func (m *Media) Validate(l *plog.Logger) error {
	if m.Name == "" {
		l.Info("Media failed validation", "reason", "empty name")
		return perr.NewErrInvalid("Media.Name cannot be empty")
	}

	return nil
}
