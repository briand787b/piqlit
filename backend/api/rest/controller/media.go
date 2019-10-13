package controller

import (
	"net/http"

	"github.com/briand787b/piqlit/core/model"
)

// MediaRequest is a request object for the Media resource
type MediaRequest struct {
	model.Media
}

// Bind does processing on the MediaRequest after it gets decoded
func (m *MediaRequest) Bind(r *http.Request) error {
	return nil
}

// MediaResponse represents the response object for Media requests
type MediaResponse struct {
	model.Media
}

// NewMediaResponse creates a new MediaResponse
func NewMediaResponse(mm *model.Media) *MediaResponse {
	return &MediaResponse{*mm}
}

// Render processes a MediaResponse before rendering in HTTP response
func (m *MediaResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// MediaResponseList represents a list of Media
type MediaResponseList struct {
	Media    []model.Media `json:"media"`
	Skip     int           `json:"skip"`
	Take     int           `json:"take"`
	NextSkip int           `json:"next_skip,omitempty"`
}

// NewMediaResponseList converts a slice of model.Media into a MediaResponseList
func NewMediaResponseList(mms []model.Media, skip, take int) *MediaResponseList {
	return &MediaResponseList{
		Media: mms,
		Skip:  skip,
		Take:  take,
	}
}

// Render does any processing ahead of the go-chi library's rendering
func (l *MediaResponseList) Render(w http.ResponseWriter, r *http.Request) error {
	if len(l.Media) >= l.Take {
		l.NextSkip = l.Skip + l.Take
	}

	return nil
}
