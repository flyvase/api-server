package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"harvest/config"
	"harvest/core/apperror"
)

type ResultImpl struct {
	Original sql.Result
}

func (r ResultImpl) LastInsertId() (int64, error) {
	return r.Original.LastInsertId()
}

type SqlImpl struct {
	Db *sql.DB
}

func NewSqlImpl() *SqlImpl {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	return &SqlImpl{Db: db}
}

func (s *SqlImpl) Exec(query string, args ...interface{}) (Result, error) {
	result, err := s.Db.Exec(query, args...)
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return nil, apperror.SqlConnClosed{Message: err.Error()}
		default:
			return nil, apperror.Unknown{Message: err.Error()}
		}
	}

	return ResultImpl{Original: result}, nil
}
