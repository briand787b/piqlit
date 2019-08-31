package model

import (
	"net"
	"time"
)

// Node represents a physical device capable of hosting and,
// potentially displaying, videos
type Node struct {
	ID        int        `sql:"int"`
	IPAddress net.IPAddr `sql:"ip_address"`
	IsMaster  bool       `sql:"is_master"`
	CreatedAt time.Time  `sql:"created_at"`
	UpdatedAt time.Time  `sql:"updated_at"`

	//
}
