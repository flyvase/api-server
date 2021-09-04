package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"harvest/config"
	"harvest/core/exception"
	"harvest/domain/repository"
)

type ResultImpl struct {
	Original sql.Result
}

func (res ResultImpl) LastInsertId() (int64, error) {
	return res.Original.LastInsertId()
}

type SqlImpl struct {
	Conn *sql.DB
}

func NewSqlRepositoryImpl() *SqlImpl {
	conn, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		panic(err)
	}

	return &SqlImpl{Conn: conn}
}

func (s SqlImpl) Exec(query string, args ...interface{}) (repository.Result, error) {
	result, err := s.Conn.Exec(query, args...)
	if err != nil {
		switch err {
		case sql.ErrConnDone:
			return nil, exception.SqlConnClosedError{Message: err.Error()}
		default:
			return nil, exception.UnknownError{Message: err.Error()}
		}
	}

	return ResultImpl{Original: result}, nil
}
