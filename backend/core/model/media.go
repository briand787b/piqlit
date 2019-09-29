package model

import (
	"time"

	"github.com/briand787b/piqlit/core/obj"
)

// Media is a container for viewable works of art, or for child media
type Media struct {
	ID           int              `sql:"id"`
	Name         string           `sql:"name"`
	Encoding     obj.Encoding     `sql:"encoding"`
	UploadStatus obj.UploadStatus `sql:"upload_status"`
	CreatedAt    time.Time        `sql:"created_at"`
	UpdatedAt    time.Time        `sql:"updated_at"`

	// non-persistence data
	ParentID *int    `sql:"parent_id"`
	children []Media `sql:"children"`
}
