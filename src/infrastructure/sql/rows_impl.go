package sql

import "database/sql"

type RowsImpl struct {
	Original *sql.Rows
}

func (r *RowsImpl) Close() error {
	return r.Original.Close()
}

func (r *RowsImpl) Err() error {
	return r.Original.Err()
}

func (r *RowsImpl) Next() bool {
	return r.Original.Next()
}

func (r *RowsImpl) Scan(dest ...interface{}) error {
	return r.Original.Scan(dest...)
}
