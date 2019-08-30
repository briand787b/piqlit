package model

import (
	"net"
	"time"
)

// Server represents a physical device capable of hosting and,
// potentially displaying, videos
type Server struct {
	ID        int        `sql:"int"`
	IPAddress net.IPAddr `sql:"ip_address"`
	IsMaster  bool       `sql:"is_master"`
	CreatedAt time.Time  `sql:"created_at"`
	UpdatedAt time.Time  `sql:"updated_at"`

	//
}
