package sql

type Sql interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
}

type Result interface {
	LastInsertId() (int64, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
	Err() error
}
