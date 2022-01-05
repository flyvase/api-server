package sql

import (
	"api-server/src/core/errors"
	"database/sql"
)

type rows struct {
	Result *sql.Rows
}

func (r *rows) Close() {
	if err := r.Result.Close(); err != nil {
		panic(err)
	}
}

func (r *rows) Next() bool {
	return r.Result.Next()
}

func (r *rows) Scan(args ...interface{}) error {
	if err := r.Result.Scan(args...); err != nil {
		return errors.Unexpected{
			Message: err.Error(),
		}
	}

	return nil
}
