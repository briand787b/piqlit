package modeltest

import (
	"context"

	"github.com/briand787b/piqlit/core/model"
)

// MediaMockStore is a mocked implementation of MediaStore
type MediaMockStore struct {
	AssociateParentIDWithChildIDsCallCount int
	AssociateParentIDWithChildIDsArgPID    []int
	AssociateParentIDWithChildIDsArgCIDs   [][]int
	AssociateParentIDWithChildIDsReturnErr []error

	DeleteByIDCallCount int
	DeleteByIDArgID     []int
	DeleteByIDReturnErr []error

	DisassociateParentIDFromChildrenCallCount int
	DisassociateParentIDFromChildrenArgPID    []int
	DisassociateParentIDFromChildrenReturnErr []error

	GetByIDCallCount   int
	GetByIDArgID       []int
	GetByIDReturnMedia []*model.Media
	GetByIDReturnErr   []error

	InsertCallCount int
	InsertArgMedia  []*model.Media
	InsertReturnErr []error

	GetByNameCallCount   int
	GetByNameArgName     []string
	GetByNameReturnMedia []*model.Media
	GetByNameReturnErr   []error

	SelectByParentIDCallCount   int
	SelectByParentIDArgPID      []int
	SelectByParentIDReturnMedia [][]model.Media
	SelectByParentIDReturnErr   []error

	UpdateCallCount int
	UpdateArgMedia  []*model.Media
	UpdateReturnErr []error
}

// AssociateParentIDWithChildIDs x
func (s *MediaMockStore) AssociateParentIDWithChildIDs(ctx context.Context, pID int, cIDs ...int) error {
	defer func() { s.AssociateParentIDWithChildIDsCallCount++ }()
	s.AssociateParentIDWithChildIDsArgPID = append(s.AssociateParentIDWithChildIDsArgPID, pID)
	s.AssociateParentIDWithChildIDsArgCIDs = append(s.AssociateParentIDWithChildIDsArgCIDs, cIDs)
	return s.AssociateParentIDWithChildIDsReturnErr[s.AssociateParentIDWithChildIDsCallCount]
}

// DeleteByID x
func (s *MediaMockStore) DeleteByID(ctx context.Context, id int) error {
	defer func() { s.DeleteByIDCallCount++ }()
	s.DeleteByIDArgID = append(s.DeleteByIDArgID, id)
	return s.DeleteByIDReturnErr[s.DeleteByIDCallCount]
}

// DisassociateParentIDFromChildren x
func (s *MediaMockStore) DisassociateParentIDFromChildren(ctx context.Context, pID int) error {
	defer func() { s.DisassociateParentIDFromChildrenCallCount++ }()
	s.DisassociateParentIDFromChildrenArgPID = append(s.DisassociateParentIDFromChildrenArgPID, pID)
	return s.DisassociateParentIDFromChildrenReturnErr[s.DisassociateParentIDFromChildrenCallCount]
}

// GetByID x
func (s *MediaMockStore) GetByID(ctx context.Context, id int) (*model.Media, error) {
	defer func() { s.GetByIDCallCount++ }()
	s.GetByIDArgID = append(s.GetByIDArgID, id)
	return s.GetByIDReturnMedia[s.GetByIDCallCount],
		s.GetByIDReturnErr[s.GetByIDCallCount]
}

// GetByName x
func (s *MediaMockStore) GetByName(ctx context.Context, name string) (*model.Media, error) {
	defer func() { s.GetByNameCallCount++ }()
	s.GetByNameArgName = append(s.GetByNameArgName, name)
	return s.GetByNameReturnMedia[s.GetByNameCallCount],
		s.GetByNameReturnErr[s.GetByNameCallCount]
}

// Insert x
func (s *MediaMockStore) Insert(ctx context.Context, m *model.Media) error {
	defer func() { s.InsertCallCount++ }()
	s.InsertArgMedia = append(s.InsertArgMedia, m)
	return s.InsertReturnErr[s.InsertCallCount]
}

// SelectByParentID x
func (s *MediaMockStore) SelectByParentID(ctx context.Context, pID int) ([]model.Media, error) {
	defer func() { s.SelectByParentIDCallCount++ }()
	s.SelectByParentIDArgPID = append(s.SelectByParentIDArgPID, pID)
	return s.SelectByParentIDReturnMedia[s.SelectByParentIDCallCount],
		s.SelectByParentIDReturnErr[s.SelectByParentIDCallCount]
}

// Update x
func (s *MediaMockStore) Update(ctx context.Context, m *model.Media) error {
	defer func() { s.UpdateCallCount++ }()
	s.UpdateArgMedia = append(s.UpdateArgMedia, m)
	return s.UpdateReturnErr[s.UpdateCallCount]
}
