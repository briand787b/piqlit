package plogtest

// MockUUIDGen is a mock that can be used for the uuidgen
// parameter for the plog.PLogger constructor
type MockUUIDGen struct {
	StringCallCount  int
	StringRetStrings []string
}

// String satisfies the fmt.Stringer interface
func (m *MockUUIDGen) String() string {
	defer func() { m.StringCallCount++ }()
	return m.StringRetStrings[m.StringCallCount]
}
