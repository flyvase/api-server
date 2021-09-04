package logger

import (
	"fmt"
	"log"

	"harvest/config"
)

func Debug(message, component string) {
	if config.Mode == "release" {
		return
	}
	e := LogEntry{
		Message:   message,
		Severity:  "DEBUG",
		Component: component,
	}

	fmt.Println(e)
}

func Info(message, component string, trace string) {
	e := LogEntry{
		Message:   message,
		Severity:  "INFO",
		Component: component,
		Trace:     trace,
	}

	fmt.Println(e)
}

func Error(
	component string, err error, trace string) {
	e := LogEntry{
		Message:   err.Error(),
		Severity:  "ERROR",
		Component: component,
		Trace:     trace,
	}

	log.Println(e)
}
