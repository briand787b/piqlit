package model

// ServerStore x
type ServerStore interface {
	GetServerByID(id int) (*Server, error)
}
