package sql

import (
	gateway "api-server/src/application/gateway/sql"
	"api-server/src/config"
	"api-server/src/core/errors"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DriverImpl struct {
	DB *sql.DB
}

func NewDriver() *DriverImpl {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	// https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &DriverImpl{DB: db}
}

func (d *DriverImpl) Exec(query string, args ...interface{}) error {
	_, err := d.DB.Exec(query, args...)
	if err != nil {
		if err == sql.ErrConnDone {
			return errors.ErrSqlConnClosed
		}

		return errors.Unexpected{
			Message: err.Error(),
		}
	}

	return nil
}

func (d *DriverImpl) Query(query string, args ...interface{}) (gateway.Rows, error) {
	rows, err := d.DB.Query(query, args...)
	if err != nil {
		if err == sql.ErrConnDone {
			return nil, errors.ErrSqlConnClosed
		}

		return nil, errors.Unexpected{
			Message: err.Error(),
		}
	}

	return &rowsImpl{
		Result: rows,
	}, nil
}

func (d *DriverImpl) QueryRow(query string, args ...interface{}) gateway.Row {
	row := d.DB.QueryRow(query, args...)
	return &rowImpl{
		Result: row,
	}
}
