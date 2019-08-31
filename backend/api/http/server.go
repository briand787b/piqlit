package phttp

import (
	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// ServerArgs holds all the necessary values for a server, but cannot
// itself serve HTTP (satisfy the http.Handler interface) unless it is
// embedded in another type.  This pattern is done to improve the maintainablity
// of this package should a new Router be substituted for Chi in the future:
// a new Server embedding ServerAbstract would be created and its ServeHTTP
// method would be called by http.ListenAndServe
type ServerArgs struct {
	plog.Logger
	model.MediaStore
	model.NodeStore
}
