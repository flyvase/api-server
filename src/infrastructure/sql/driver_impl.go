package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"harvest/src/config"
	"harvest/src/core/apperror"
)

type DriverImpl struct {
	DB *sql.DB
}

func NewDriverImpl() *DriverImpl {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	return &DriverImpl{DB: db}
}

func (d *DriverImpl) Exec(query string, args ...interface{}) error {
	_, err := d.DB.Exec(query, args...)
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return apperror.SqlConnClosed{Message: err.Error()}
		default:
			return apperror.Unknown{Message: err.Error()}
		}
	}

	return nil
}

func (d *DriverImpl) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := d.DB.Query(query, args...)
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return nil, apperror.SqlConnClosed{Message: err.Error()}
		default:
			return nil, apperror.Unknown{Message: err.Error()}
		}
	}

	return &RowsImpl{Original: rows}, nil
}

func (d *DriverImpl) QueryRow(query string, args ...interface{}) (Row, error) {
	row := d.DB.QueryRow(query, args...)
	err := row.Err()
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return nil, apperror.SqlConnClosed{Message: err.Error()}
		default:
			return nil, apperror.Unknown{Message: err.Error()}
		}
	}

	return &RowImpl{Original: row}, nil
}
