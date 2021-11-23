package logger

import "log"

func Error(
	component string,
	err error,
	trace string,
) {
	e := LogEntry{
		Message:   err.Error(),
		Severity:  "ERROR",
		Trace:     trace,
		Component: component,
	}

	log.Println(e)
}
