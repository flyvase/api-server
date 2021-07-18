package logger

import (
	"context"
	"fmt"

	"harvest/config"
)

const (
	debug = "DEBUG"
	info  = "INFO"
	err   = "ERROR"
)

func Debug(message, component string) {
	if config.Mode == "release" {
		return
	}
	e := LogEntry{
		Message:   message,
		Severity:  debug,
		Component: component,
	}

	fmt.Println(e)
}

func Info(ctx context.Context, message, component string) {
	e := LogEntry{
		Message:   message,
		Severity:  info,
		Component: component,
		Trace:     ctx.Value("trace").(string),
	}

	fmt.Println(e)
}

func Error(ctx context.Context, message, component string) {
	e := LogEntry{
		Message:   message,
		Severity:  err,
		Component: component,
		Trace:     ctx.Value("trace").(string),
	}

	fmt.Println(e)
}
