package sql

import (
	"database/sql"
	"harvest/src/config"
)

type DriverImpl struct {
	DB *sql.DB
}

func NewDriver() *DriverImpl {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	return &DriverImpl{DB: db}
}

func (d *DriverImpl) Exec(query string, args ...interface{}) error {
	_, err := d.DB.Exec(query, args...)
	return err
}

func (d *DriverImpl) Query(query string, args ...interface{}) (Rows, error) {
	return d.DB.Query(query, args...)
}

func (d *DriverImpl) QueryRow(query string, args ...interface{}) Row {
	return d.DB.QueryRow(query, args...)
}
