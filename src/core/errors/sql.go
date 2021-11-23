package errors

import "errors"

var (
	ErrSqlConnClosed = errors.New("sql connection is already closed")
	ErrSqlNoRows     = errors.New("no rows in sql result set")
)
