package sql

import "database/sql"

type rowsImpl struct {
	Result *sql.Rows
}

func (r *rowsImpl) Close() error {
	return r.Result.Close()
}

func (r *rowsImpl) Next() bool {
	return r.Result.Next()
}

func (r *rowsImpl) Scan(args ...interface{}) error {
	return r.Result.Scan(args...)
}
