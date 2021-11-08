package apperror

type SqlConnClosed struct {
	Message string
}

func (e SqlConnClosed) Error() string {
	return e.Message
}

type EmptySqlResult struct {
	Message string
}

func (e EmptySqlResult) Error() string {
	return e.Message
}
