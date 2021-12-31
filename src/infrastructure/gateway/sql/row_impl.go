package sql

import (
	"api-server/src/core/errors"
	"database/sql"
)

type rowImpl struct {
	Result *sql.Row
}

func (r *rowImpl) Scan(args ...interface{}) error {
	err := r.Result.Scan(args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrSqlNoRows
		}

		return errors.Unexpected{
			Message: err.Error(),
		}
	}

	return nil
}
