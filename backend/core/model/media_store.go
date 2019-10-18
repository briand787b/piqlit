package model

import (
	"context"
)

// MediaTxCtlStore x
type MediaTxCtlStore interface {
	Begin(context.Context) (MediaTxCtlStore, error)
	Commit(context.Context) error
	Rollback(context.Context) error
	MediaStore
}

// MediaStore is anything that can store and retrieve Media records from a database
type MediaStore interface {
	AssociateParentIDWithChildIDs(ctx context.Context, pID int, cIDs ...int) error
	DeleteByID(ctx context.Context, id int) error
	DisassociateParentIDFromChildren(ctx context.Context, pID int) error
	GetByID(ctx context.Context, id int) (*Media, error)
	GetByName(ctx context.Context, name string) (*Media, error)
	Insert(ctx context.Context, m *Media) error
	SelectByParentID(ctx context.Context, pID int) ([]Media, error)
	Update(ctx context.Context, m *Media) error
}
