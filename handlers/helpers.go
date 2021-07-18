package handlers

import (
	"fmt"
	"strings"

	"harvest/config"
)

func GetTraceId(header string) string {
	parts := strings.Split(header, "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", config.ProjectId, parts[0])
	}
	return ""
}
