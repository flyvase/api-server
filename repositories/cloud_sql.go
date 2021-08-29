package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"harvest/config"
)

func InitMySqlConnection() (*sql.DB, error) {
	dbPool, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
