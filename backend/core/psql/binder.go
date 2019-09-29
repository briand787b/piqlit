package psql

// Binder is an interface for something which can bind queries (Tx, DB)
type binder interface {
	DriverName() string
	Rebind(string) string
	BindNamed(string, interface{}) (string, []interface{}, error)
}
