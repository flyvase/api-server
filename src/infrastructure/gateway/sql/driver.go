package sql

import (
	"api-server/src/config"
	"api-server/src/core/errors"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	Db *sql.DB
}

func NewDriver() *Driver {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	// https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Driver{Db: db}
}

func (d *Driver) Exec(query string, args ...interface{}) error {
	_, err := d.Db.Exec(query, args...)
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

func (d *Driver) Query(query string, args ...interface{}) (*rows, error) {
	r, err := d.Db.Query(query, args...)
	if err != nil {
		if err == sql.ErrConnDone {
			return nil, errors.ErrSqlConnClosed
		}

		return nil, errors.Unexpected{
			Message: err.Error(),
		}
	}

	return &rows{
		Result: r,
	}, nil
}

func (d *Driver) QueryRow(query string, args ...interface{}) *row {
	r := d.Db.QueryRow(query, args...)
	return &row{
		Result: r,
	}
}
