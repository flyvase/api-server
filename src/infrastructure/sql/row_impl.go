package sql

import (
	"database/sql"
	"harvest/src/core/apperror"
)

type RowImpl struct {
	Original *sql.Row
}

func (r *RowImpl) Scan(dest ...interface{}) error {
	if err := r.Original.Scan(dest...); err != nil {
		if err == sql.ErrNoRows {
			return apperror.EmptySqlResult{Message: err.Error()}
		} else {
			return apperror.Unknown{Message: err.Error()}
		}
	}

	return nil
}

func (r *RowImpl) Err() error {
	return r.Original.Err()
}
