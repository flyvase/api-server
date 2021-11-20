package sql

type Rows interface {
	Close() error
	Next() bool
	Scan(...interface{}) error
}
