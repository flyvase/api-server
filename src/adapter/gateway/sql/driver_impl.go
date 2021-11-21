package sql

import (
	"database/sql"
	gateway "harvest/src/application/gateway/sql"
	"harvest/src/config"
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
	return err
}

func (d *DriverImpl) Query(query string, args ...interface{}) (gateway.Rows, error) {
	return d.DB.Query(query, args...)
}

func (d *DriverImpl) QueryRow(query string, args ...interface{}) gateway.Row {
	return d.DB.QueryRow(query, args...)
}
