package logger

import (
	"context"
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

func Info(ctx context.Context, message, component string) {
	e := LogEntry{
		Message:   message,
		Severity:  "INFO",
		Component: component,
		Trace:     ctx.Value("trace").(string),
	}

	fmt.Println(e)
}

func Error(ctx context.Context, component string, err error) {
	msg := fmt.Sprintf("%+v", err)
	e := LogEntry{
		Message:   msg,
		Severity:  "ERROR",
		Component: component,
		Trace:     ctx.Value("trace").(string),
	}

	log.Println(e)
}
