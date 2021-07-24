package core

type DSConnErr struct {
	Original error
}

func (e DSConnErr) Error() string {
	return e.Original.Error()
}

type UnknownErr struct {
	Original error
}

func (e UnknownErr) Error() string {
	return e.Original.Error()
}
