package errors

type Unexpected struct {
	Message string
}

func (e *Unexpected) Error() string {
	return e.Message
}
