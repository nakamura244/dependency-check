package database

// Rows is interface
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
