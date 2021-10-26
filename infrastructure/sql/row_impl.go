package sql

import "database/sql"

type RowImpl struct {
	Original *sql.Row
}

func (r *RowImpl) Scan(dest ...interface{}) error {
	return r.Original.Scan(dest...)
}

func (r *RowImpl) Err() error {
	return r.Original.Err()
}
