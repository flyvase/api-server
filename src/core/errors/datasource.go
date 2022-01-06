package errors

import "errors"

var (
	ErrDatasourceConnClosed = errors.New("datasource connection is already closed")
	ErrDataNotFound         = errors.New("data not found")
)
