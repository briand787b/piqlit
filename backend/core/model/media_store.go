package model

// MediaStore x
type MediaStore interface {
	GetMediaByID(id int) (*Media, error)
}
