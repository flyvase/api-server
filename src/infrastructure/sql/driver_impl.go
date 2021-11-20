package sql

import "database/sql"

type DriverImpl struct {
	DB *sql.DB
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
