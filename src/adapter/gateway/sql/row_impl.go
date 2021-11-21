package sql

import "database/sql"

type RowImpl struct {
	Result *sql.Row
}

func (r *RowImpl) Scan(args ...interface{}) error {
	return r.Result.Scan(args...)
}
