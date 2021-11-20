package sql

import "database/sql"

type RowsImpl struct {
	Result *sql.Rows
}

func (r *RowsImpl) Close() error {
	return r.Result.Close()
}

func (r *RowsImpl) Next() bool {
	return r.Result.Next()
}

func (r *RowsImpl) Scan(args ...interface{}) error {
	return r.Result.Scan(args...)
}
