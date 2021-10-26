package sql

type Rows interface {
	Close() error
	Err() error
	Next() bool
	Scan(...interface{}) error
}
