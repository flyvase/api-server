package apperror

type Unknown struct {
	Message string
}

func (e Unknown) Error() string {
	return e.Message
}
