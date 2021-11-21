package sql

type Driver interface {
	Exec(string, ...interface{}) error
	Query(string, ...interface{}) (Rows, error)
	QueryRow(string, ...interface{}) Row
}
