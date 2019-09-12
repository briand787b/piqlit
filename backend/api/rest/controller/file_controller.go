package controller

import (
	"encoding/json"
	"net/http"

	"github.com/briand787b/piqlit/core/obj"
	"github.com/sirupsen/logrus"
)

// FileController controls File paths
type FileController struct {
	l  *logrus.Logger
	os obj.ObjectStore
}

// List returns all the files managed by piqlit on this node
func (fc *FileController) List(w http.ResponseWriter, r *http.Request) {
	fs, err := fc.os.ListNames(r.Context())
	if err != nil {
		fc.l.Error(err)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(fs); err != nil {
		fc.l.Error(err)
		w.WriteHeader(500)
	}
}
