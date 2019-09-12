package controller

import (
	"net/http"
)

// FileResponseList presents a list of model.File
type FileResponseList struct {
	Files []string
}

// NewFileResponseList instantiates a new FileResponseList
func NewFileResponseList(fs []string) *FileResponseList {
	return &FileResponseList{
		Files: fs,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (frl *FileResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
