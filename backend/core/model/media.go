package model

import (
	"time"
)

// Media is a container for viewable works of art, or for child media
type Media struct {
	ID            int       `sql:"id"`
	Title         string    `sql:"title"`
	ThumbnailName string    `sql:"thumbnail_name"`
	ReleaseDate   time.Time `sql:"release_date"`
	ParentID      *int      `sql:"parent_id"`
	CreatedAt     time.Time `sql:"created_at"`
	UpdatedAt     time.Time `sql:"updated_at"`

	// non-persistence data
	childMedia []Media
	servers    []Node
}
