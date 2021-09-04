package exception

type SqlConnClosedError struct {
	Message string
}

func (e SqlConnClosedError) Error() string {
	return e.Message
}
