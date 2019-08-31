package model

// NodeStore x
type NodeStore interface {
	GetNodeByID(id int) (*Node, error)
}
