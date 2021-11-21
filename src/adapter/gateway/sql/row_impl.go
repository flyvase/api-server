package sql

import "database/sql"

type rowImpl struct {
	Result *sql.Row
}

func (r *rowImpl) Scan(args ...interface{}) error {
	return r.Result.Scan(args...)
}
