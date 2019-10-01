package model

import "context"

// MediaStore is anything that can store and retrieve Media records
// from a database
type MediaStore interface {
	DeleteByID(ctx context.Context, id int) error
	FindByParentID(ctx context.Context, pID int) ([]Media, error)
	GetByID(ctx context.Context, id int) (*Media, error)
	Insert(ctx context.Context, m *Media) error
}
