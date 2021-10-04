package sql

type Sql interface {
	Exec(string, ...interface{}) (Result, error)
}

type Result interface {
	LastInsertId() (int64, error)
}
