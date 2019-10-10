package model

import (
	"context"
)

// MediaStore is anything that can store and retrieve Media records from a database
type MediaStore interface {
	AssociateParentIDWithChildIDs(ctx context.Context, pID int, cIDs ...int) error
	DeleteByID(ctx context.Context, id int) error
	DisassociateParentIDFromChildIDs(ctx context.Context, pID int, cIDs ...int) error
	GetByID(ctx context.Context, id int) (*Media, error)
	Insert(ctx context.Context, m *Media) error
	SelectByName(ctx context.Context, name string) ([]Media, error)
	SelectByParentID(ctx context.Context, pID int) ([]Media, error)
	Update(ctx context.Context, m *Media) error
}
