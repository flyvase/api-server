package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"harvest/config"
)

func InitMySqlConnection() (*sql.DB, error) {
	dbPool, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dbPool, nil
}
