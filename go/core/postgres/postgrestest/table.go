package postgrestest

// Table enumerates a Postgres table
type Table int

const (
	// Media represents the Media table
	Media Table = iota

	// ChildMedia represents a child record
	// in the Media table
	ChildMedia

	// GrandchildMedia represents a grandchild
	// record in the Media table
	GrandchildMedia
)
