package exception

type UnknownError struct {
	Message string
}

func (e UnknownError) Error() string {
	return e.Message
}
