package sql

import "database/sql"

type Sql interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
}

type Result interface {
	LastInsertId() (int64, error)
}

// type Row interface {
// 	Hoge()
// }
