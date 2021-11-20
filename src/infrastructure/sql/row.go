package sql

type Row interface {
	Scan(...interface{}) error
}
