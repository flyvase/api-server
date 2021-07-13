package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type LogEntry struct {
	Message   string `json:"message"`
	Severity  string `json:"severity,omitempty"`
	Trace     string `json:"logging.googleapis.com/trace,omitempty"`
	Component string `json:"component,omitempty"`
}

func (e LogEntry) String() string {
	if e.Severity == "" {
		e.Severity = "INFO"
	}
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v\n", err)
	}
	return string(out)
}

const (
	info    = "INFO"
	warning = "WARNING"
	err     = "ERROR"
)

func GetTraceId(header string) string {
	parts := strings.Split(header, "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", ProjectId, parts[0])
	}
	return ""
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

func Warning(ctx context.Context, message, component string) {
	e := LogEntry{
		Message:   message,
		Severity:  warning,
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
