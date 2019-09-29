package model

// MediaStore is anything that can store and retrieve Media records
// from a database
type MediaStore interface {
	GetByID(id int) (*Media, error)
}
