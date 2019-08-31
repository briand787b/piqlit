package model

// NodePGStore is a NodeStore backed by Postgres
type NodePGStore struct {
}

// GetNodeByID retrieves a Node by its ID
func (nps *NodePGStore) GetNodeByID(id int) (*Node, error) {
	return nil, nil
}
