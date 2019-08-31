package phttp

import (
	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// ServerArgs aggregates all the server arguments
type ServerArgs struct {
	plog.Logger
	model.MediaStore
	model.ServerStore
}
