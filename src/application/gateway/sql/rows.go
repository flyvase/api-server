package sql

type Rows interface {
	Close()
	Next() bool
	Scan(...interface{}) error
}
