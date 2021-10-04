package apperror

type SqlConnClosed struct {
	Message string
}

func (e SqlConnClosed) Error() string {
	return e.Message
}
