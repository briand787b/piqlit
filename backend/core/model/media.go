package model

import (
	"time"

	"github.com/briand787b/piqlit/core/obj"
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
